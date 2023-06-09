package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

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

	wg.Wait()
	fmt.Println("Main goroutine stopped.")
}
