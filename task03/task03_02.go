package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println(solutionQ(numbers))

	fmt.Println(time.Since(startTime))
}

func solutionQ(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num * num
	}

	return result
}
