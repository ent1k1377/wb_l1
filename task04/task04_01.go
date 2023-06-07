package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Надо подробнее рассмотреть

func worker(workerID int, dataChannel <-chan string, wg *sync.WaitGroup) {
	for c := range dataChannel {
		time.Sleep(time.Second * 3)
		fmt.Printf("workerID: %d, message: %s\n", workerID, c)
	}
	fmt.Println(123)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	fmt.Print("Введите число воркеров: ")
	var numberWorker int
	_, err := fmt.Scan(&numberWorker)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	wg.Add(numberWorker)
	ch := make(chan string, numberWorker)

	for i := 0; i < numberWorker; i++ {
		go worker(i, ch, &wg)
	}

	go func() {
		for {
			fmt.Print("You message: ")
			var input string
			fmt.Scan(&input)
			ch <- input
		}
	}()

	sigCh := make(chan os.Signal, 1)
	go signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	close(ch)
	wg.Wait()
	fmt.Println("exit")
}
