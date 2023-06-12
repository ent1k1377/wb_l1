package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100
	fmt.Println(cap(v), len(v))
	v = append(v, b)
	fmt.Println(cap(v), len(v))
}

func main() {
	a := make([]int8, 0, 6)
	for i := int8(0); i < 5; i++ {
		a = append(a, i)
	}
	fmt.Println(a, cap(a), len(a))
	someAction(a, 6)
	fmt.Println(a)
}
