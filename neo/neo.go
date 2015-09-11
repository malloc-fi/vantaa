package neo

import (
	"os"

	"github.com/jmcvetta/neoism"
)

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
