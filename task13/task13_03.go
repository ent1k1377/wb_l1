package main

import "fmt"

func main() {
	a := 33
	b := 11

	fmt.Printf("Before: a = %d, b = %d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After: a = %d, b = %d", a, b)
}
