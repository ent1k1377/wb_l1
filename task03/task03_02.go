package main

import (
	"fmt"
	"time"
)

func main() {

	var numbers []int
	func(numbers *[]int) {
		for i := 100000; i < 1000000; i += 2 {
			*numbers = append(*numbers, i)
		}
	}(&numbers)

	startTime := time.Now()

	result := 0
	for _, num := range numbers {
		result += num * num
	}

	fmt.Println(result)

	fmt.Println(time.Since(startTime), time.Since(startTime).Nanoseconds(), "ns")
}
