package vantaadb

import (
	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/settings"
)

var db *neoism.Database = &neoism.Database{}

// Init sets db with a valid neoism.Database value based on the environment's
// DbUrl setting. Panic if unable to connect to the database
func Init() {
	var err error = nil
	db, err = neoism.Connect(settings.Get().DbUrl)
	if err != nil {
		panic(err)
	}
}

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

// Connect create return a neoism.Database instance that can be used to process
// cypher queries through the REST API
func Connect() *neoism.Database {
	if db.Url == "" {
		Init()
	}
	return db
}
