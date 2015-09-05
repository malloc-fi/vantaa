package testhelper

import (
	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/neo"
)

func ClearNeo() {
	db := neo.Connect()
	cq := neoism.CypherQuery{
		Statement: `MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE n,r`,
	}
	if err := db.Cypher(&cq); err != nil {
		panic(err)
	}
}
