package main

import (
	"errors"
	"fmt"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	err := removeElement(&numbers, 2)
	fmt.Println(numbers)
	if err != nil {
		fmt.Println(err)
	}
}

// removeElement удаляет элемент с указанным индексом из среза чисел.
// В случае успешного удаления, срез будет изменен.
// Если индекс невалидный, возвращается ошибка.
func removeElement(numbers *[]int, index int) error {
	if index < 0 || index > len(*numbers)-1 {
		return errors.New("неверный индекс")
	}
	*numbers = append((*numbers)[:index], (*numbers)[index+1:]...)
	return nil
}
