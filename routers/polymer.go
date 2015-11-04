package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// SetPolymerAppRoutes sets the path for polymer app
func SetPolymerAppRoutes(router *mux.Router) *mux.Router {

	adminAppFiles := http.FileServer(http.Dir("./admin/app/"))
	//	adminBowerFiles := http.FileServer(http.Dir("./admin/bower_components/"))

	router.PathPrefix("/admin/").Handler(http.StripPrefix("/admin/", adminAppFiles))
	//	router.PathPrefix("/admin/bower_components/").Handler(http.StripPrefix("/admin/bower_components/", adminBowerFiles))

	return router
}
