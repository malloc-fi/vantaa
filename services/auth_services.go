package services

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nathandao/vantaa/api/parameters"
	"github.com/nathandao/vantaa/core/authentication"
	"github.com/nathandao/vantaa/services/models/user"
)

func Login(requestUser *user.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerareToken(requestUser.ID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(paramenters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}

func ReFreshToken(request *user.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.ID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req,
		func(token *jwt.Token) (interface{}, error) {
			return authBackend.PublicKey, nil
		})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
