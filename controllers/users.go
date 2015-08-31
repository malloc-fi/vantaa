package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nathandao/vantaa/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type (
	UserController struct {
		Session *mgo.Session
	}
)

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func (uc UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func (uc UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
