package main

import (
	"fmt"
	"sync"
)

func main() {
	var numbers []int

	// Анонимная функция, которая заполняет слайс numbers значениями от 65 до 122
	func(list *[]int) {
		for i := 65; i < 123; i++ {
			*list = append(*list, i)
		}
	}(&numbers)

	var cm sync.Map // Создание экземпляра sync.Map, которая представляет конкурентно-безопасное отображение

	var wg sync.WaitGroup
	wg.Add(len(numbers))

	// Запуск горутин для каждого значения из среза numbers
	for _, v := range numbers {
		go func(v int) {
			defer wg.Done()

			cm.Store(v, string(rune(v)))              // Сохранение значения в sync.Map
			fmt.Printf("%d:%s\n", v, string(rune(v))) // Вывод значения
		}(v)
	}

	wg.Wait() // Ожидание завершения всех горутин

	lenCM := 0
	cm.Range(func(k, v any) bool {
		lenCM += 1
		return true
	})

	// Проверка, что количество элементов в sync.Map равно количеству элементов в срезе numbers
	if lenCM == len(numbers) {
		fmt.Println("\nit's all good")
	}
}
