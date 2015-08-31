package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct{}

const (
	Dbname = "vantaa-local"
	Dbuser = "guynathan"
	Dbpass = "Nothing123"
	Dburl  = "ds055812.mongolab.com:55812"
)

func (d *DB) Session() *mgo.Session {
	s, err := mgo.Dial("mongodb://" + Dbuser + ":" + Dbpass + "@" + Db.url + "/" + Dbname)
	if err != nil {
		panic(err)
	}
	return s
}
