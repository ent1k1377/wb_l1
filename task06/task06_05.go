package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopped.")
				return
			default:
				// Ваш код выполнения задачи
				fmt.Println("Doing some work...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Main goroutine stopped.")
}
