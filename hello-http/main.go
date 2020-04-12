package main

import (
	"fmt"
	"net/http"

	"github.com/alex-parra/go-learn/hello-http/controllers"
)

func main() {
	port := 4549

	fmt.Println("Hello HTTP")
	fmt.Println("Listening on port ", port)

	controllers.RegisterControllers()
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
