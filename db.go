package main

import (
	"log"

	"github.com/cayleygraph/cayley"
	_ "github.com/cayleygraph/cayley/graph/bolt"
)

// TODO: Add check bd sessions
// var gErr = graph.InitQuadStore("bolt", "/vantaa/db", nil)
var store, dbErr = cayley.NewGraph("bolt", "./tmp/db", nil)

func Db() (*cayley.Handle, error) {
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	return store, dbErr
}
