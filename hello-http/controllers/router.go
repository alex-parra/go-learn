package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers routes
func RegisterControllers() {
	uc := newUserController()
	http.Handle("/", *uc)
}

func sendJSON(w io.Writer, data interface{}) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
