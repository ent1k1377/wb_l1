package main

import (
	"fmt"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(numbers))
	resultCh := make(chan int, len(numbers))

	for _, num := range numbers {
		go func(num int) {
			defer wg.Done()
			square := num * num
			resultCh <- square
		}(num)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	sum := 0
	for ch := range resultCh {
		sum += ch
	}
	fmt.Println(sum)
	fmt.Println(time.Since(startTime), time.Since(startTime).Nanoseconds(), "ns")
}
