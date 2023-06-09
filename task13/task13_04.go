package main

import "fmt"

func main() {
	a := 2
	b := 77

	fmt.Printf("Before: a = %d, b = %d\n", a, b)
	a = a * b
	b = a / b
	a = a / b
	fmt.Printf("After: a = %d, b = %d", a, b)
}
