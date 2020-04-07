package controllers

import (
	"fmt"
	"lenslocked.com/views"
	"net/http"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
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
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "name is:", form.Name)
	fmt.Fprintln(w, "email is:", form.Email)
	fmt.Fprintln(w, "password is:", form.Password)
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
