package models

import (
	"github.com/nathandao/vantaa/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Content struct {
	Id          bson.ObjectId   `json:"id" bson:"_id"`
	Title       string          `json:"title" bson:"title"`
	Body        string          `json:"body" bson:"body"`
	MainImage   bson.ObjectId   `json:"main_image" bson:"main_image"`
	Categories  []bson.ObjectId `json:"categories" bson:"categories"`
	Slug        string          `json:"slug" bson:"slug"`
	Created     time.Time       `json:"created" bson:"created"`
	PublishDate time.Time       `json:"publish_date" bson:"publish_date"`
	Status      string          `json:"status" bson:"status"`
	Type        string          `json:"type" bson:"type"`
}

func (c *Content) GetAuthor() User {

}
