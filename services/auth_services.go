package services

import (
	"encoding/json"
	"net/http"

	// jwt "github.com/dgrijalva/jwt-go"
	"github.com/nathandao/vantaa/core/auth"
	"github.com/nathandao/vantaa/services/models/user"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func Login(requestUser *user.User) (int, []byte) {
	authBackend := auth.InitJwtAuthBackend()
	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.Id)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(TokenAuthentication{token})
			return http.StatusOK, response
		}
	}
	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *user.User) []byte {
	authBackend := auth.InitJwtAuthBackend()
	token, err := authBackend.GenerateToken(requestUser.Id)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

// func Logout(req *http.Request) error {
// 	authBackend := auth.InitJwtAuthBackend()
// 	tokenRequest, err := jwt.ParseFromRequest(req,
// 		func(token *jwt.Token) (interface{}, error) {
// 			return authBackend.PublicKey, nil
// 		})
// 	if err != nil {
// 		return err
// 	}
// 	tokenString := req.Header.Get("Authorization")
// 	return authBackend.Logout(tokenString, tokenRequest)
// }
