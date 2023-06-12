package main

import (
	"fmt"
	"strings"
)

type data struct {
	input  string
	result bool
}

func main() {
	testTable := []data{
		{
			input:  "abcd",
			result: true,
		},
		{
			input:  "abCdefAaf",
			result: false,
		},
		{
			input:  "aabcd",
			result: false,
		},
	}

	// Проходим по тестовой таблице и проверяем соответствие результатов функции
	// areAllCharactersUnique ожидаемым значениям.
	for _, d := range testTable {
		fmt.Println(areAllCharactersUnique(d.input) == d.result)
	}
}

// areAllCharactersUnique проверяет, содержит ли входная строка только уникальные символы.
// Возвращает true, если все символы в строке уникальны, и false в противном случае.
func areAllCharactersUnique(input string) bool {
	// Создаем карту для отслеживания уникальных символов.
	characters := make(map[rune]bool)
	// Приводим входную строку к нижнему регистру для регистронезависимой проверки.
	lowerInput := strings.ToLower(input)

	// Проходим по каждому символу в строке.
	for _, c := range lowerInput {
		// Если символ уже присутствует в карте, значит он не уникален, и возвращаем false.
		if _, exists := characters[c]; exists {
			return false
		}
		// Добавляем символ в карту как уникальный.
		characters[c] = true
	}
	// Если все символы прошли проверку и не было повторов, возвращаем true.
	return true
}
