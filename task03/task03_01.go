package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем слайс длинны n
	numbers := []int{2, 4, 6, 8, 10}

	// Создаем WaitGroup для синхронизации горутин и добавляем в нее длину слайса
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	// Создаем буферизованный канал для передачи квадратов чисел
	squareCh := make(chan int, len(numbers))

	// Запускаем горутины для вычисления квадратов чисел
	for _, num := range numbers {
		// Создание n горутин
		go func(num int) {
			defer wg.Done()
			square := num * num
			// Передаем в канал квадрат чисел
			squareCh <- square
		}(num)
	}

	// Запускаем горутину для закрытия канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(squareCh)
	}()

	// Суммируем квадраты чисел из канала
	sum := 0
	for ch := range squareCh {
		sum += ch
	}

	// Выводим результат
	fmt.Println(sum)
}
