package main

import "fmt"

func adder(base int) func(x int) int {
	sum := base
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	const start = 2
	sum := adder(start)
	for i := start; i < 25; i++ {
		fmt.Println(sum(i))
	}
}
