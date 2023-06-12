package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Остановка программы через селект и контекст с таймаутом
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Создание контекста с таймаутом 5 секунд
	defer cancel()                                                          // Отмена контекста при выходе из функции main

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				// Остановка выполнения горутины по сигналу из контекста при истечении времени таймаута
				fmt.Println("stop go func")
				return
			default:
				// Бесконечный цикл с печатью "..."
				fmt.Println("...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	wg.Wait()                // Ожидание завершения работы горутины
	fmt.Println("main stop") // Вывод сообщения о завершении работы программы
}
