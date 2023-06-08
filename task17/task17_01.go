package main

import (
	"errors"
	"fmt"
)

type data struct {
	numbers []int
	target  int
	result  int
}

func main() {
	testTable := []data{
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			2,
			0,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			4,
			1,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			6,
			2,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			8,
			3,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			12,
			5,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			18,
			8,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			20,
			9,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			21,
			-1,
		},
		{
			[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
			1,
			-1,
		},
	}
	for _, d := range testTable {
		index, err := binarySearch(d.numbers, d.target)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(index == d.result)
	}
}

func binarySearch(numbers []int, target int) (int, error) {
	l, r := 0, len(numbers)-1
	for l <= r {
		m := (l + r) / 2
		if numbers[m] == target {
			return m, nil
		} else if numbers[m] > target {
			r = m - 1
		} else if numbers[m] < target {
			l = m + 1
		}
	}
	return -1, errors.New("number does not exist")
}
