package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nathandao/vantaa/controllers"
	"github.com/nathandao/vantaa/core/auth"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/api/auth/token", controllers.Login).Methods("POST")

	router.Handle("/api/auth/refresh-token", negroni.New(
		negroni.HandlerFunc(auth.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.RefreshToken),
	)).Methods("GET")

	router.Handle("/api/auth/logout",
		negroni.New(
			negroni.HandlerFunc(auth.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	return router
}
