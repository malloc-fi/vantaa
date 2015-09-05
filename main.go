package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"

	"github.com/jmcvetta/neoism"
	"github.com/julienschmidt/httprouter"
	"github.com/nathandao/vantaa/model"
)

type templateHandler struct {
	once     sync.Once
	filename string
	template *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		log.Println("loaded template:", filepath.Join("templates", t.filename))
		t.template = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.template.Execute(w, nil)
}

func init() {
	u := model.User{
		Name:     "admin",
		Email:    "admin@example.com",
		Password: "password",
	}

	if u2, _ := model.FindUser(neoism.Props{"name": "admin"}); u2 == nil {
		u.Save()
	}
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := &templateHandler{filename: "loginform.html"}
	t.ServeHTTP(w, r)
}

func Admin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")
	password := r.FormValue("password")

	if name == "" || password == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	passwordd, _ := model.GenrerateSalt([]byte(password))
	u, _ := model.FindUser(neoism.Props{"name": name})
	log.Println(bytes.Compare(passwordd, u.PasswordDigest))

	if u == nil || bytes.Compare(passwordd, u.PasswordDigest) != 0 {
		http.Redirect(w, r, "/login", http.StatusFound)
	}

	t := &templateHandler{filename: "admin.html"}
	t.ServeHTTP(w, r)
}

func main() {
	r := httprouter.New()

	r.GET("/login", Login)
	r.GET("/admin", Admin)

	http.ListenAndServe(":9290", r)
}
