package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	increment int            // Переменная для инкремента счетчика
	mu        sync.RWMutex   // RWMutex для синхронизации доступа к счетчику
	wg        sync.WaitGroup // WaitGroup для ожидания завершения горутин
}

func main() {
	numberIteration := 1000 // Количество итераций
	var counter Counter     // Создание экземпляра структуры Counter

	counter.wg.Add(numberIteration) // Установка счетчика ожидания
	for i := 0; i < numberIteration; i++ {
		go func() {
			defer counter.wg.Done() // Уменьшение счетчика ожидания при завершении горутины

			counter.mu.Lock()   // Блокировка мьютекса для обеспечения эксклюзивного доступа к счетчику
			counter.increment++ // Инкрементация счетчика
			counter.mu.Unlock() // Разблокировка мьютекса
		}()
	}

	counter.wg.Wait()              // Ожидание завершения всех горутин
	fmt.Println(counter.increment) // Вывод значения счетчика
}
