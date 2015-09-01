package main

import (
	"fmt"
	"github.com/nathandao/vantaa/db"
	"github.com/nathandao/vantaa/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	s := db.Session()
	defer s.Close()
	s.SetSafe(&mgo.Safe{})

	c := s.DB(db.Dbname).C("Users")

	result := models.User{}
	c.Find(bson.M{"name": "nathandao"}).One(&result)

	u := models.User{
		Name:     "nathandao",
		Email:    "nathan@guynathan.com",
		Password: "Nothing123",
	}

	u.Save(s)

	fmt.Println(us.Name)
	fmt.Println("Hello World")
}
