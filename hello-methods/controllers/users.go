package controllers

import (
	"net/http"
	"regexp"
)

// UserController controls
type UserController struct {
	userIDPattern *regexp.Regexp
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (uc UserController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Users...."))
}
