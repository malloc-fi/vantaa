package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nathandao/vantaa/services"
	"github.com/nathandao/vantaa/services/models/user"
)

func Login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(user.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	fmt.Println(requestUser)
	fmt.Println(r.FormValue("email"))

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(user.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
