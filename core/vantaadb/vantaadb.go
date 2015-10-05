package vantaadb

import (
	"os"

	"github.com/jmcvetta/neoism"
)

// CypherQuery is basically neoism.CypherQuery. Adding here so we don't have
// to include neoism separately in other packages, which should provide clearer
// abtraction layers.
type CypherQuery interface{}

// PropString construct a neois.Props using the prefix - which is the Type's
// name and its parameters, and returns a map[string]interface{} that can be
// used as the cypher query's parameters.
func PropString(prefix string, props neoism.Props) string {
	qstr := ""
	for k, _ := range props {
		if k != "id" {
			qstr += prefix + "." + k + " = {" + k + "} and "
		} else {
			qstr += "id(" + prefix + ") = {" + k + "} and "
		}
	}
	// remove trailing and
	if qstr != "" {
		qstr = qstr[:len(qstr)-5]
	}

	return qstr
}

// Connect create a new neoism.Database instance that can be used to send
// cypher queries through the REST API
func Connect() *neoism.Database {
	url := os.Getenv("NEO4J")
	if url == "" {
		url = "http://neo4j:foobar@localhost:7474"
	}
	db, err := neoism.Connect(url)
	if err != nil {
		panic(err)
	}
	return db
}
