package main

import (
	"fmt"
	"github.com/nathandao/vantaa/neo"
)

func main() {
	n := neo.Connect()
	fmt.Println(n.Url)
}
