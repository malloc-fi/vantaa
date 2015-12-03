package main

import (
	"net/http"

	"github.com/nathandao/vantaa/Godeps/_workspace/src/github.com/codegangsta/negroni"
	"github.com/nathandao/vantaa/routers"
	"github.com/rs/cors"
	//	"github.com/nathandao/vantaa/services/models/user"
)

// func init() {
// 	user.CreateDummyUser()
// }

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Content-type", "X-CSRF-TOKEN", "Content-Length", "Authorization"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
	})

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	router := routers.InitRoutes()

	n.Use(c)
	n.UseHandler(router)

	http.ListenAndServe(":9292", n)
}
