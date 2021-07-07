package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome website</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>To get in touch, contact <a href=\"mailto:support@lenslock.com\">support@lenslock.com</a></p>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<p>Where to contact <a href=\"mailto:support@lenslock.com\">support@lenslock.com</a></p><p>Address: 555 main st., Owen City, CA 99999</p>")
}

func noFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Sorry we can not find the page</h1><p> If it continue to happen, contact <a href=\"mailto:support@lenslock.com\">support@lenslock.com</a></p>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(noFound)
	_ = http.ListenAndServe(":3000", r)
}
