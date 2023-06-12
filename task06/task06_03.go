package main

import (
	"fmt"
	"sync"
	"time"
)

// Остановка программы через селект и закрытие канала
func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{}) // Канал типа struct{} для отправки сигнала остановки горутины

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stop:
				// Остановка выполнения горутины по сигналу из канала stop
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
	close(stop)                 // Закрытие канала stop для отправки сигнала остановки горутины

	wg.Wait()                // Ожидание завершения работы горутины
	fmt.Println("main stop") // Вывод сообщения о завершении работы программы
}
