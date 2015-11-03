package auth

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request,
	next http.HandlerFunc) {

	authBackend, err := InitJwtAuthBackend()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	token, err := jwt.ParseFromRequest(
		req,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			} else {
				return authBackend.PublicKey, nil
			}
		},
	)

	if err == nil && token.Valid && !authBackend.IsTerminated(req.Header.Get("Authorization")) {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
