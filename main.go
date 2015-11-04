package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/nathandao/vantaa/routers"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
