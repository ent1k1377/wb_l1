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

func removeElement(numbers *[]int, index int) error {
	if index < 0 || index > len(*numbers)-1 {
		return errors.New("invalid index")
	}
	*numbers = append((*numbers)[:index], (*numbers)[index+1:]...)
	return nil
}
