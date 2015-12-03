package auth

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func AllowCrossOrigin(w http.ResponseWriter, req *http.Request,
	next http.HandlerFunc) {
	fmt.Println("cross origin called!")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	next(w, req)
}

func RequireTokenAuthentication(w http.ResponseWriter, req *http.Request,
	next http.HandlerFunc) {
	fmt.Println("require token called!")

	if req.Method == "OPTIONS" {
		return
	}

	// if origin := req.Header.Get("Origin"); origin != "" {
	// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// 	w.Header().Set("Access-Control-Allow-Headers",
	// 		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// }

	authBackend, err := InitJwtAuthBackend()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}

	fmt.Println(req.Header)

	token, err := jwt.ParseFromRequest(
		req,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf(
					"Unexpected signing method: %v",
					token.Header["alg"],
				)
			} else {
				return authBackend.PublicKey, nil
			}
		},
	)

	if err == nil &&
		token.Valid &&
		!authBackend.IsTerminated(req.Header.Get("Authorization")) {
		fmt.Println("PASSED")
		next(w, req)
	} else {
		fmt.Println(req.Header.Get("Authorization"))
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func ValidateTokenAuthentication(w http.ResponseWriter, req *http.Request,
	next http.HandlerFunc) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	authBackend, err := InitJwtAuthBackend()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	token, err := jwt.ParseFromRequest(
		req,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf(
					"Unexpected signing method: %v",
					token.Header["alg"],
				)
			} else {
				return authBackend.PublicKey, nil
			}
		},
	)

	if err == nil && token.Valid && !authBackend.IsTerminated(
		req.Header.Get("Authorization")) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
