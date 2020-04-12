package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/alex-parra/go-learn/hello-http/models"
)

// UserController controls
type UserController struct {
	Path          string
	userIDPattern *regexp.Regexp
}

func newUserController() *UserController {
	return &UserController{
		Path:          "/users",
		userIDPattern: regexp.MustCompile(`^/(\d+)/?`),
	}
}

func (uc UserController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path == uc.Path {
		switch req.Method {
		case http.MethodGet:
			uc.getAll(res, req)
		case http.MethodPost:
			uc.createUser(res, req)
		default:
			res.WriteHeader(http.StatusNotImplemented)
		}
	} else {

	}
}

func (uc *UserController) stub(data string, res http.ResponseWriter) {
	sendJSON(res, data)
}

func (uc *UserController) getAll(res http.ResponseWriter, req *http.Request) {
	sendJSON(res, models.GetUsers())
}

func (uc *UserController) get(id int, res http.ResponseWriter, req *http.Request) {
	u, err := models.GetUserByID(id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
	}
	sendJSON(res, fmt.Sprintf("models.GetByID(%d)", u.ID))
}

func (uc *UserController) createUser(res http.ResponseWriter, req *http.Request) {
	u, err := uc.parseRequest(req)
	if err != nil {
		res.WriteHeader(http.StatusNotAcceptable)
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(err.Error()))
		return
	}
	sendJSON(res, u)
}

func (uc *UserController) parseRequest(req *http.Request) (models.User, error) {
	dec := json.NewDecoder(req.Body)
	var u models.User
	err := dec.Decode(&u)
	if err == nil {
		return models.User{}, err
	}
	return u, nil
}
