package routers

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/nathandao/vantaa/controllers"
	"github.com/nathandao/vantaa/core/auth"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {

	tokenRouter := mux.NewRouter()

	tokenRouter.Handle("/api/auth/token/new", negroni.New(
		negroni.HandlerFunc(controllers.Login),
	)).Methods("POST")

	tokenRouter.Handle("/api/auth/logout", negroni.New(
		negroni.HandlerFunc(auth.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Logout),
	))

	tokenRouter.Handle("/api/auth/token/refresh", negroni.New(
		negroni.HandlerFunc(auth.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.RefreshToken),
	)).Methods("GET")

	tokenRouter.Handle("/api/auth/token/validate", negroni.New(
		negroni.HandlerFunc(auth.RequireTokenAuthentication),
		negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
			w.WriteHeader(http.StatusOK)
		}),
	))

	router.PathPrefix("/api/auth").Handler(negroni.New(
		negroni.Wrap(tokenRouter),
	))

	return router
}
