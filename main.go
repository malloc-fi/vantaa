package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/nathandao/vantaa/routers"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)
	n.UseHandler(router)

	http.ListenAndServe(":9292", n)
}
