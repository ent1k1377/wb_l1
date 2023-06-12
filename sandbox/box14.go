package main

import "fmt"

func main() {
	slice := make([]string, 0, 3)
	slice = append(slice, "a", "a")
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
