package main

import "fmt"

func main() {
	fmt.Println("\nFor loop to completion")
	var i int
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\nFor loop 3 terms")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var x int
	fmt.Println("\nFor loop infinite")
	for {
		if x == 5 {
			break
		}
		fmt.Println(x)
		x++
	}

	fmt.Println("\nLoop collections")
	fruits := []string{"Apple", "Banana", "Kiwi"}
	for i, v := range fruits {
		fmt.Println(i, v)
	}

	fmt.Println("")
	statuses := map[string]int{"OK": 200, "Unauthorized": 401, "Not Found": 404, "Server Error": 500}
	for desc, code := range statuses {
		fmt.Println(code, ":", desc)
	}
}
