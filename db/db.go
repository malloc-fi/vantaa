package db

import (
	"github.com/jmcvetta/neoism"
	"os"
)

func Connect() *neoism.Database {
	url := os.Getenv("NEO4J")
	if url == "" {
		url = "http://localhost:7474/data/db"
	}
	db, err := neoism.Connect(url)
	if err != nil {
		panic(err)
	}
	return db
}
