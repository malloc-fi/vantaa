package models

import (
	"errors"
	"github.com/nathandao/vantaa/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Model interface {
	Find(params bson.M)
	Create(resource interface{})
	Delete(params bson.M)
}
