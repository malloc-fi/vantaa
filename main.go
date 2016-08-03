package main

import (
	"fmt"
)

func main() {
	u := User{
		Name:     "testuser",
		Email:    "user@example.com",
		Password: "somepassword",
	}
	u.Create()

	newU, _ := GetObjects("", "")
	fmt.Println(newU)
	// Create a brand new graph
	// Initialize the database
	// store, err := Db()

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Quad(subject, predicate, object, label interface{})
	// Triple(subject, pridcate, label interface{})

	// triple := cayley.Triple("Kevin", "Bacon", "son of")
	// store.AddQuad(triple)

	// triple2 := cayley.Triple("pharse of the day", "Kevin", "son of")
	// store.AddQuad(triple2)

	// store.AddQuad(quad.Make("phrase of the day", "is of course", "Hello World!", "demo graph"))
	// store.AddQuad(quad.Make("user", "HAS_POST", "Post 1", "posts graph"))
	// store.AddQuad(quad.Make("user", "HAS_POST", "Post 2", "posts graph"))
	// store.AddQuad(quad.Make("user", "HAS_POST", "Post 3", "posts graph"))
	// store.AddQuad(quad.Make("user", "HAS_POST", "Post 4", "posts graph"))

	// // Now we create the path, to get to our data
	// p := cayley.StartPath(store, quad.String("user")).Out(quad.String("HAS_POST"))

	// // Now we get an iterator for the path (and optimize it, the second return is if it was optimized,
	// // but we don"t care for now)
	// it, _ := p.BuildIterator().Optimize()
	// // remember to cleanup after yourself
	// defer it.Close()

	// // Now for each time we can go to next iterator
	// nxt := graph.AsNexter(it)
	// // remember to cleanup after yourself
	// defer nxt.Close()

	// // While we have items
	// for nxt.Next() {
	// 	token := it.Result()                // get a ref to a node
	// 	value := store.NameOf(token)        // get the value in the node
	// 	nativeValue := quad.NativeOf(value) // this converts nquad values to normal Go type

	// 	fmt.Println(nativeValue) // print it!
	// }
	// if err := nxt.Err(); err != nil {
	// 	log.Fatalln(err)
	// }
}
