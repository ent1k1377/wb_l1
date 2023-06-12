package main

import (
	"fmt"
	"strings"
)

type data struct {
	input  string
	result string
}

func main() {
	testTable := []data{
		{
			"snow dog sun",
			"sun dog snow",
		},
		{
			"golang postgres redis nats",
			"nats redis postgres golang",
		},
		{
			"moscow berlin tokio",
			"tokio berlin moscow",
		},
	}

	for _, d := range testTable {
		fmt.Println(reverseWordsInSentence(d.input) == d.result)
	}
}

func reverseWordsInSentence(sentence string) string {
	sentenceSplit := strings.Split(sentence, " ") // Разбиение предложения на отдельные слова с помощью разделителя " "
	if len(sentenceSplit) == 0 {
		return sentence // Если предложение пустое, возвращается исходное предложение без изменений
	}

	newSentence := sentenceSplit[0] // Начальное слово
	for _, w := range sentenceSplit[1:] {
		newSentence = w + " " + newSentence // Добавление слова перед начальным словом с пробелом
	}
	return newSentence // Возвращение предложения с перевернутым порядком слов
}
