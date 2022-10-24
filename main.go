package main

import (
	"fmt"
	"strconv"
)

func main() {
	rand := "22e718e3e1065a4f0e8991324664bbb8781dbf449836414586a73465e8d0926d"
	fmt.Println("Hello world!")
	x := add(3, 13)

	fmt.Println("The answer is", x)

	if strconv.Itoa(x) == rand {
		fmt.Println("wot?")
	}
}

// Please ignore this
func add(a int, b int) int {
	fmt.Printf("Performing %d + %d\n", a, b)
	return a + b
}
