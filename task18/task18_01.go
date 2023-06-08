package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	increment int
	mu        sync.RWMutex
	wg        sync.WaitGroup
}

func main() {
	numberIteration := 1000
	var counter Counter

	counter.wg.Add(numberIteration)
	for i := 0; i < numberIteration; i++ {
		go func() {
			defer counter.wg.Done()

			counter.mu.Lock()
			counter.increment++
			counter.mu.Unlock()
		}()
	}

	counter.wg.Wait()
	fmt.Println(counter.increment)
}
