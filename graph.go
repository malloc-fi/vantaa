package main

import (
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
)

// Vantaa blog name
const VANTAA_BLOG = "vantaa blog"
const HAS_USER = "has user"

// Types of graphs
const USERS_GRAPH = "users graph"
const POSTS_GRAPH = "posts graph"

// Relationships
const HAS_NAME = "has name"
const HAS_EMAIL = "has email"
const HAS_PASSWORD_DIGEST = "has password digest"

// GetObjects retrieves all objects that relates to the subject through
// predicate.
func GetObjects(subject, predicate string) ([]interface{}, error) {
	nativeValues := []interface{}{}

	// Now we create the path, to get to our data
	p := cayley.StartPath(store, quad.String(subject)).Out(quad.String(predicate))

	// Now we get an iterator for the path (and optimize it, the second return
	// is if it was optimized, but we don"t care for now)
	it, _ := p.BuildIterator().Optimize()
	defer it.Close()

	// Now for each time we can go to next iterator
	nxt := graph.AsNexter(it)
	defer nxt.Close()

	// While we have items
	for nxt.Next() {
		token := it.Result()
		value := store.NameOf(token)
		nativeValues = append(nativeValues, quad.NativeOf(value))
	}

	return nativeValues, nxt.Err()
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
