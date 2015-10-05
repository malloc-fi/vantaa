package neo4jstore

import (
	"github.com/jmcvetta/neoism"
)

func Cleanup() {
	db, _ := neoism.Connect("http://neo4j:foobar@localhost:7474")
	cq := neoism.CypherQuery{
		Statement: `MATCH (n) OPTIONAL MATCH (n)-[r]-() DELETE n,r`,
	}
	if err := db.Cypher(&cq); err != nil {
		panic(err)
	}
}
