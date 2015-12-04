package auth

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func RequireTokenAuthentication(w http.ResponseWriter, req *http.Request,
	next http.HandlerFunc) {

	// Preflight handling
	if req.Method == "OPTIONS" {
		return
	}

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

	if err == nil &&
		token.Valid &&
		!authBackend.IsTerminated(req.Header.Get("Authorization")) {
		next(w, req)
	} else {
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
