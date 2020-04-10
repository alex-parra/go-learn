package main

import (
	"fmt"

	"github.com/alex-parra/go-learn/hello-collections/models"
)

func main() {

	arraysAndSlices()
	maps()
	structs()

}

func arraysAndSlices() {
	var alex [3]string
	alex[0] = "Alex"
	alex[1] = "Parra"
	alex[2] = "Silva"
	fmt.Println(alex)

	filipa := [3]string{"Filipa", "Carvalho", "Melo"}
	fmt.Println(filipa)

	sebastian := []string{"Sebastião", "Melo", "Parra"}
	fmt.Println(sebastian)

	familyName := alex[2:]
	fmt.Println(familyName)
}

func maps() {
	family := map[string]string{"father": "alex", "mother": "filipa", "son": "Sebastião"}
	fmt.Println(family)

	family["twins"] = "Vicente & Inês"
	fmt.Println(family)
}

func structs() {
	var alex models.User
	alex.FirstName = "Alexandre"
	alex.LastName = "Parra"
	alex.Age = 39

	filipa := models.User{
		FirstName: "Filipa",
		LastName:  "Melo",
		Age:       40,
	}

	fmt.Println(alex)
	fmt.Println(filipa)
}
