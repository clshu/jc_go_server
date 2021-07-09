package controllers

import (
	"fmt"
	"net/http"

	"github.com/clshu/jc_go_server/views"
)

// NewUser is used to create a new Users controller.
// It panics if the thempates are not parsed correctly.
func NewUser() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}


type Users struct {
	NewView *views.View
}

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users Create")
}
