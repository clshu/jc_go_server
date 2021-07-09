package main

import (
	"net/http"

	"github.com/clshu/jc_go_server/controllers"
	"github.com/gorilla/mux"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUser()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
