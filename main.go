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

	fs := http.FileServer(http.Dir("./admin/app"))
	fs2 := http.FileServer(http.Dir("./admin/bower_components"))
	http.Handle("/admin/", http.StripPrefix("/admin/", fs))
	http.Handle("/admin/bower_components", http.StripPrefix("/admin/bower_components", fs2))
	http.ListenAndServe(":9292", nil)
}
