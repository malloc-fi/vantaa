package testhelpers

import (
	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/core/vantaadb"
)

func ClearDb() {
	db := vantaadb.Connect()
	cq := neoism.CypherQuery{
		Statement: `MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE n,r`,
	}
	if err := db.Cypher(&cq); err != nil {
		panic(err)
	}
}
