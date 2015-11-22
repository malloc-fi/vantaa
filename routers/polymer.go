package routers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var env = os.Getenv("GO_ENV")

// SetPolymerAppRoutes sets the path for polymer app
func SetPolymerAppRoutes(router *mux.Router) *mux.Router {

	router = router.StrictSlash(true)

	if env == "production" {
		adminAppFiles := http.FileServer(http.Dir("./admin/dist/"))
		router.PathPrefix("/admin/").Handler(http.StripPrefix("/admin/", adminAppFiles))
	} else {
		adminAppFiles := http.FileServer(http.Dir("./admin/app/"))
		adminBowerFiles := http.FileServer(http.Dir("./admin/bower_components/"))

		router.PathPrefix("/admin/bower_components/").Handler(http.StripPrefix("/admin/bower_components/", adminBowerFiles))
		router.PathPrefix("/admin").Handler(http.StripPrefix("/admin", adminAppFiles))
		router.PathPrefix("/admin/").Handler(http.StripPrefix("/admin/", adminAppFiles))
	}

	return router
}
