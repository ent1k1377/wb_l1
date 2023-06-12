package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Остановка программы через селект и контекст
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				// Остановка выполнения горутины по сигналу от контекста
				fmt.Println("go func stop")
				return
			default:
				// Бесконечный цикл с печатью "..."
				fmt.Println("...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second) // Ожидание 5 секунд
	cancel()                    // Отмена контекста

	wg.Wait()                // Ожидание завершения работы горутины
	fmt.Println("main stop") // Вывод сообщения о завершении работы программы
}
