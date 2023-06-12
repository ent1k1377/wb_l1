package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	sum := 0

	// WaitGroup для синхронизации горутин
	var wg sync.WaitGroup
	// Mutex для обеспечения безопасного доступа к переменной sum из разных горутин
	var mu sync.Mutex

	// Устанавливаем счетчик WaitGroup равным количеству чисел
	wg.Add(len(numbers))
	// Запускаем горутины для вычисления квадратов чисел
	for _, num := range numbers {
		go func(num int) {
			// Захватываем мьютекс перед обновлением переменной sum
			mu.Lock()
			// Освобождаем мьютекс после обновления переменной sum
			defer mu.Unlock()
			// Вычисляем квадрат числа и добавляем его к sum
			sum += num * num
			// Уменьшаем счетчик WaitGroup после завершения горутины
			wg.Done()
		}(num)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	// Выводим сумму квадратов чисел
	fmt.Println(sum)
}
