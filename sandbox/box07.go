package main

import "fmt"

func main() {
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		m[i] = i
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
