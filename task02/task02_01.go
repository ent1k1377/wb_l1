package main

import (
	"fmt"
	"sync"
)

// 1. Main. Создаем слайс чисел, создаем waitGroup и добавляем в нее длину слайса,
// в for range перебираем последовательно все числа и запускаем горутины (одна горутина рассчитывает квадрат переданного числа и выводит в Stdout)
// Мы используем WaitGroup для ожидания завершения группы задач, через wg.Add мы увеличиваем внутренний счетчик wg
// через wg.Done мы уменьшаем счетчик на один (defer указывает что вызов функции произойдет перед завершением горутины),
// а wg.Wait блокирует main и ждет когда в счетчике будет 0

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(numbers))

	for _, n := range numbers {
		go func(n int) {
			defer wg.Done()
			square := n * n
			fmt.Printf("%d * %d = %d\n", n, n, square)
		}(n)
	}

	wg.Wait()
	fmt.Println("exit")
}
