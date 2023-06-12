package main

import (
	"fmt"
)

type data struct {
	input  string
	result string
}

func main() {
	testTable := []data{
		{
			input:  "главрыба",
			result: "абырвалг",
		},
		{
			input:  "wildberies",
			result: "seirebdliw",
		},
		{
			input:  "task19_01.go",
			result: "og.10_91ksat",
		},
	}

	for _, d := range testTable {
		fmt.Println(reverseString(d.input) == d.result)
	}
}

func reverseString(s string) string {
	runes := []rune(s) // Преобразование строки в срез рун

	for i := 0; i < len(runes)/2; i++ {
		swapIndex := len(runes) - i - 1                         // Индекс для обмена символов
		runes[i], runes[swapIndex] = runes[swapIndex], runes[i] // Обмен символов
	}
	return string(runes) // Преобразование среза рун обратно в строку
}
