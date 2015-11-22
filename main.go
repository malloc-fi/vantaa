package main

import (
	"net/http"

	"github.com/nathandao/vantaa/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"github.com/nathandao/vantaa/routers"
	//	"github.com/nathandao/vantaa/services/models/user"
)

// func init() {
// 	user.CreateDummyUser()
// }

func main() {
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	router := routers.InitRoutes()
	n.UseHandler(router)

	http.ListenAndServe(":9292", n)
}
