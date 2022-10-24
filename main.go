package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	x := add(3, 13)

	fmt.Println("The answer is", x)
}

func add(a int, b int) int {
	fmt.Printf("Performing %d + %d\n", a, b)
	return a + b
}
