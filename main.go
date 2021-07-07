package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to my awesome website</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<p>To get in touch, contact <a href=\"mailto:support@lenslock.com\">support@lenslock.com</a></p>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Sorry we can not find the page</h1><p>contact us if you continue having the problem</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	_ = http.ListenAndServe(":3000", nil)
}
