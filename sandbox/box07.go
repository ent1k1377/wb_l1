package main

import "fmt"

func main() {
	m := make(map[int]int, 5)
	m[0] = 1
	m[1] = 124
	m[2] = 281
	m[3] = 333
	m[4] = 444
	for k, v := range m {
		fmt.Println(k, v)
	}
}
