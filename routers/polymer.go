package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetPolymerAppRoutes sets the path for polymer app
func SetPolymerAppRoutes(router *mux.Router) *mux.Router {
	router.Handle("/", http.FileServer(http.Dir(("./public/front")))).Methods("GET")
	router.Handle("/admin", http.FileServer(http.Dir("./public/admin"))).Methods("GET")

	return router
}
