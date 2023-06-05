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
	for _, d := range testTable {
		fmt.Println(areAllCharactersUnique(d.input) == d.result)
	}
}

func areAllCharactersUnique(input string) bool {
	// Описание
	characters := make(map[rune]bool)
	lowerInput := strings.ToLower(input)

	for _, c := range lowerInput {
		if _, exists := characters[c]; exists {
			return false
		}
		characters[c] = true
	}
	return true
}
