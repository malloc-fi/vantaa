package neo

import (
	"github.com/jmcvetta/neoism"
	"os"
)

func Connect() *neoism.Database {
	url := os.Getenv("NEO4J")
	if url == "" {
		url = "http://neo4j:admin@localhost:7474"
	}
	db, err := neoism.Connect(url)
	if err != nil {
		panic(err)
	}
	return db
}
