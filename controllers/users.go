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

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form = SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Email: "+form.Email)
	fmt.Fprintln(w, "Password: "+form.Password)
}
