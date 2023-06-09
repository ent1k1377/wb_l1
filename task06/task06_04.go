package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stop := false

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			if stop {
				fmt.Println("Goroutine stopped.")
				return
			}

			fmt.Println("Doing some work...")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Stopping goroutine...")
	stop = true

	wg.Wait()
	fmt.Println("Main goroutine stopped.")
}
