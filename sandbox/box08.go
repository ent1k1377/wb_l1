package main

import "fmt"

type User struct {
	Name string
}

func main() {
	m1 := make([]int, 0)
	m2 := make([]int, 1)

	n1 := new(int)
	n2 := new([]int)
	n3 := new(chan int)
	n4 := new(User)
	(*n4).Name = "qwe"
	*n1 = 2
	fmt.Println(m1, m2, *n1, *n2, n3, *n4)
}
