package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	numbers := []int{2, 4, 6, 8, 10}
	solution(numbers)

	fmt.Println(time.Since(startTime))
}

func solution(numbers []int) {
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for i, n := range numbers {
		go func(index int, num int) {
			defer wg.Done()
			square := num * num
			fmt.Printf("%d * %d = %d\n", num, num, square)
		}(i, n)
	}

	wg.Wait()
}
