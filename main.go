package main

import (
	"github.com/google/cayley"
	"github.com/google/cayley/graph/bolt"

	"fmt"
)

func main() {
	bolt.init()
	store, err := cayley.NewGraph("bolt", "/Users/nathan/vantaa.db", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(store)
}
