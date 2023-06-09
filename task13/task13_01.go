package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Printf("Before: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("After: a = %d, b = %d", a, b)
}
