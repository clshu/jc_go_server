package main

import (
	"net/http"

	"github.com/clshu/jc_go_server/views"
	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Render(w, nil); err != nil {
		panic(err)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := signupView.Render(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("boostrap", "views/home.gohtml")
	contactView = views.NewView("boostrap", "views/contact.gohtml")
	signupView = views.NewView("boostrap", "views/signup.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/signup", signup)

	_ = http.ListenAndServe(":3000", r)
}
