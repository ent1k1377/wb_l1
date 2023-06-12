package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stop := false // Флаг остановки выполнения горутины

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			if stop {
				// Остановка выполнения горутины, если флаг stop установлен в true
				fmt.Println("stop go func")
				return
			}

			fmt.Println("...")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(5 * time.Second) // Ожидание 5 секунд
	stop = true                 // Установка флага stop в true для остановки горутины

	wg.Wait()                // Ожидание завершения работы горутины
	fmt.Println("main stop") // Вывод сообщения о завершении работы программы
}
