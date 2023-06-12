package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем слайс чисел и waitGroup
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	// Мы используем WaitGroup для ожидания завершения группы задач, через wg.Add мы увеличиваем внутренний счетчик wg
	wg.Add(len(numbers))

	// в for range перебираем последовательно все числа и запускаем горутины (одна горутина рассчитывает квадрат переданного числа и выводит в Stdout)
	for _, n := range numbers {
		go func(n int) {
			// через wg.Done мы уменьшаем счетчик на один (defer указывает что вызов функции произойдет перед завершением горутины),
			defer wg.Done()
			square := n * n
			fmt.Printf("%d * %d = %d\n", n, n, square)
		}(n)
	}
	// wg.Wait блокирует main и ждет когда в счетчике будет 0
	wg.Wait()
	fmt.Println("exit")
}
