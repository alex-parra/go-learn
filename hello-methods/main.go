package main

import (
	"fmt"
	"net/http"

	"github.com/alex-parra/go-learn/hello-methods/controllers"
)

func main() {
	fmt.Println("Hello Methods")

	controllers.RegisterControllers()
	http.ListenAndServe(":4549", nil)
}
