package main

import (
	"fmt"

	"github.com/alex-parra/go-learn/hello-strings/strutil"
)

func main() {
	var name string = "Alex"
	s := strutil.Reverse("hello")
	s = s + ` ` + name // you cannot change string values by index, but you can get values instead.
	fmt.Printf("%s\n", s)

	fruits := []string{"Banana", "Maça", "Kiwi"}
	fruits = append(fruits, "Uva")
	fmt.Println(fruits[3])
}
