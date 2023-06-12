package main

import (
	"fmt"
	"sync"
	"time"
)

// Остановка программы через селект и канал
func main() {
	var wg sync.WaitGroup
	done := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-done:
				// Остановка выполнения горутины
				fmt.Println("stop go func")
				return
			default:
				// Бесконечный цикл с печатью "..."
				fmt.Println("...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second) // Ожидание 5 секунд
	done <- true                // Отправка сигнала остановки в канал done

	wg.Wait()                // Ожидание завершения работы горутины
	fmt.Println("main stop") // Вывод сообщения о завершении работы программы
}
