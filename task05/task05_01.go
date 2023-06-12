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

	timeout := make(chan struct{}, 1)
	go func() {
		// Горутина для установки таймаута работы программы
		time.Sleep(time.Second * time.Duration(duration))
		timeout <- struct{}{}
	}()

	message := make(chan string, 1)
	go func() {
		// Горутина для обработки сообщений
		for c := range message {
			fmt.Printf("message: %s\n", c)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		// Горутина для чтения ввода пользователя
		for scanner.Scan() {
			input := scanner.Text()
			message <- input
		}
	}()

	<-timeout // Ожидание получения сигнала таймаута

	close(message)        // Закрытие канала message
	fmt.Println("\nexit") // Вывод сообщения о завершении программы
}
