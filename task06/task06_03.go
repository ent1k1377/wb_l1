package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stop:
				fmt.Println("Goroutine stopped.")
				return
			default:
				// Ваш код выполнения задачи
				fmt.Println("Doing some work...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Stopping goroutine...")
	close(stop)

	wg.Wait()
	fmt.Println("Main goroutine stopped.")
}
