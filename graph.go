package main

import (
	"fmt"
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
)

// Vantaa blog name
const VANTAA_BLOG = "vantaa blog"
const HAS_USER = "has user"

// Relationships
const HAS_NAME = "has name"
const HAS_EMAIL = "has email"
const HAS_PASSWORD_DIGEST = "has password digest"

// GetObjects retrieves all objects that relates to the subject through
// predicate.
func GetObjects(subject, predicate string) ([]interface{}, error) {
	store, _ := Db()
	nativeValues := []interface{}{}

	// Now we create the path, to get to our data
	p := cayley.StartPath(store, quad.String(subject)).Out(quad.String(predicate))

	// Now we get an iterator for the path (and optimize it, the second return
	// is if it was optimized, but we don"t care for now)
	it, _ := p.BuildIterator().Optimize()
	defer it.Close()

	// While we have items
	for it.Next() {
		token := it.Result()
		value := store.NameOf(token)
		fmt.Println("SEarching value:", value)
		nativeValues = append(nativeValues, quad.NativeOf(value))
	}

	return nativeValues, it.Err()
}

// NewTransaction shorthand.
func NewTransaction() *graph.Transaction {
	return cayley.NewTransaction()
}

// MakeQuad shorthand.
func MakeQuad(subject, predicate, object, graph string) quad.Quad {
	return quad.Make(subject, predicate, object, graph)
}

// ApplyTransactions shorthand.
func ApplyTransaction(t *graph.Transaction) error {
	store, err := Db()
	if err != nil {
		return err
	}
	err = store.ApplyTransaction(t)

	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// RemoveAllQuads removes all quads in graph.
func RemoveAllQuads() error {
	store, err := Db()

	if err != nil {
		return err
	}

	it := store.QuadsAllIterator()
	it, _ = it.Optimize()
	defer it.Close()

	for it.Next() {
		store.RemoveQuad(store.Quad(it.Result()))
	}

	return it.Err()
}

// RemoveAllNodes removes all nodes in graph.
func RemoveAllNodes() error {
	store, err := Db()

	if err != nil {
		return err
	}
	it := store.NodesAllIterator()
	defer it.Close()

	for it.Next() {
		store.RemoveNode(it.Result())
	}

	return it.Err()
}
