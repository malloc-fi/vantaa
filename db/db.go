package db

import (
	"gopkg.in/mgo.v2"
)

const (
	Dbname = "vantaa-local"
	Dbuser = "guynathan"
	Dbpass = "Nothing123"
	Dburl  = "ds055812.mongolab.com:55812"
)

func Session() *mgo.Session {
	dbu := "mongodb://" + Dbuser + ":" + Dbpass + "@" + Dburl + "/" + Dbname
	s, err := mgo.Dial(dbu)
	if err != nil {
		panic(err)
	}
	return s
}
