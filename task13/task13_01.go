package main

import "fmt"

func main() {
	a := 12.0000000000000000000000000000000001
	b := -5.1
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}

func swap01(n1 *int, n2 *int) {

}
