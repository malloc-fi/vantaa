package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nathandao/vantaa/controllers"
	"github.com/nathandao/vantaa/core/auth"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {

	tokenRouter := mux.NewRouter()

	tokenRouter.Handle("/api/auth/refresh-token", negroni.New(
		negroni.HandlerFunc(auth.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.RefreshToken),
	)).Methods("GET")

	tokenRouter.Handle("/api/auth/logout", negroni.New(
		negroni.HandlerFunc(auth.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Logout),
	))

	tokenRouter.Handle("/api/auth/token", negroni.New(
		negroni.HandlerFunc(controllers.Login),
	)).Methods("POST")

	router.PathPrefix("/api/auth").Handler(negroni.New(
		negroni.Wrap(tokenRouter),
	))

	return router
}
