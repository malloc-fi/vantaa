package controller

import (
	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/neo"
)

type UserController struct {
	Db *neoism.Database
}

func NewUserController() UserController {
	return UserController{Db: neo.Connect()}
}
