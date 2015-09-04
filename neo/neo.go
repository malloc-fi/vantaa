package neo

import (
	"os"

	"github.com/jmcvetta/neoism"
)

func PropString(props neoism.Props) string {
	qstr := ""
	for k, _ := range props {
		qstr += "\"" + k + "\": {" + k + "},"
	}

	// remove trailing comma
	if qstr != "" {
		qstr = qstr[:len(qstr)-1]
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
