package main

import (
	"fmt"
)

type data2 struct {
	numbers []int
	result  []int
}

func main() {
	testTable := []data2{
		{
			[]int{10, 5, 2, 1, 4},
			[]int{1, 2, 4, 5, 10},
		},
		{
			[]int{7, 2, 1, 6, 8, 5, 3, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{-10, 1234, 90, 90, -12},
			[]int{-12, -10, 90, 90, 1234},
		},
	}
	for _, v := range testTable {
		fmt.Printf("Before: %v\n", v.numbers)
		quickSort2(&v.numbers)
		fmt.Printf("After: %v\n", v.numbers)
		fmt.Println(EqualSlice2(v.result, v.numbers))
	}
}

func quickSort2(q *[]int) {
	numbers := *q
	if len(numbers) <= 1 {
		return
	}

	supportElement := numbers[len(numbers)-1]
	i := 0
	for j := 0; j < len(numbers)-1; j++ {
		if numbers[j] <= supportElement {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
		}
	}

	numbers[i], numbers[len(numbers)-1] = numbers[len(numbers)-1], numbers[i]
	l := numbers[:i]
	r := numbers[i+1:]

	quickSort2(&l)
	quickSort2(&r)
}

func EqualSlice2(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
