package controllers

import "net/http"

// RegisterControllers routes
func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
