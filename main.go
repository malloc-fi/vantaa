package main

import (
	"github.com/google/cayley"

	"fmt"
)

func main() {
	store, err := cayley.NewGraph("bolt", "/Users/nathan/vantaa.db", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(store)
}
