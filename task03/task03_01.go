package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println(solution(numbers))

	fmt.Println(time.Since(startTime))
}

func solution(numbers []int) int {
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	result := 0
	for _, num := range numbers {
		go func(num int) {
			defer wg.Done()
			result += num * num
		}(num)
	}
	wg.Wait()
	return result
}
