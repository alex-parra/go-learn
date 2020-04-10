package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max, _ := strconv.ParseInt(os.Args[1], 10, 64)
	var i int64
	for i = 1; i < max; i++ {
		if i%5 == 0 || i%7 == 0 {
			if i%5 == 0 {
				fmt.Print("Ping")
			}
			if i%7 == 0 {
				fmt.Print("Pong")
			}
		} else {
			fmt.Print(i)
		}
		fmt.Println("")
	}

}
