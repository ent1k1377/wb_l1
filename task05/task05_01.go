package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Print("Введите количество времени в секундах для работы программы: ")
	var duration int
	_, err := fmt.Scan(&duration)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * time.Duration(duration))
		timeout <- true
	}()

	ch := make(chan string, 1)
	go func() {
		for c := range ch {
			time.Sleep(time.Second)
			fmt.Printf("message: %s\n", c)
		}
		fmt.Println(len(ch), cap(ch))
	}()

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			input := scanner.Text()
			ch <- input
		}
	}()

	<-timeout
	close(ch)
	fmt.Println("exit")
}
