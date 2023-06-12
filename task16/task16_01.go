package main

import (
	"fmt"
)

type data01 struct {
	numbers []int
	result  []int
}

func main() {
	// Тестовые данные для сортировки
	testTable := []data01{
		{
			[]int{10, 5, 2, 1, 4},
			[]int{1, 2, 4, 5, 10},
		},
		{
			[]int{7, 2, 1, 6, 8, 5, 3, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			[]int{-10, 1234, 90, 90, -12},
			[]int{-12, -10, 90, 90, 1234},
		},
	}

	// Проход по всем тестовым данным
	for _, v := range testTable {
		fmt.Printf("Before: %v\n", v.numbers)
		quickSort01(v.numbers)
		fmt.Printf("After: %v\n", v.numbers)
		fmt.Println("Сортировка прошла успешно: ", EqualSlice01(v.result, v.numbers), "\n")
	}
}

func quickSort01(numbers []int) {
	// Базовый случай - если длина среза меньше или равна 1, нет необходимости сортировать
	if len(numbers) <= 1 {
		return
	}

	// Опорный элемент - последний элемент в срезе
	supportElement := numbers[len(numbers)-1]

	// Индекс для разделения на две части
	i := 0
	for j := 0; j < len(numbers)-1; j++ {
		if numbers[j] <= supportElement {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
		}
	}

	// Помещаем опорный элемент в правильную позицию
	numbers[i], numbers[len(numbers)-1] = numbers[len(numbers)-1], numbers[i]

	// Рекурсивно сортируем две части среза
	quickSort01(numbers[:i])
	quickSort01(numbers[i+1:])
}

func EqualSlice01(s1, s2 []int) bool {
	// Проверка на равенство длин срезов
	if len(s1) != len(s2) {
		return false
	}

	// Проверка на равенство элементов срезов по индексам
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
