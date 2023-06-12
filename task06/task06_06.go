package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Остановка программы через селект и остановку через Ctrl+C
func main() {
	var wg sync.WaitGroup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM) // Регистрация сигналов SIGINT и SIGTERM для канала stop

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-stop:
				// Остановка выполнения горутины по сигналу из канала stop при получении сигнала SIGINT или SIGTERM
				fmt.Println("\nstop go func")
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
