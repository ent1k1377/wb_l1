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
	sentenceSplit := strings.Split(sentence, " ")
	if len(sentenceSplit) == 0 {
		return sentence
	}

	newSentence := sentenceSplit[0]
	for _, w := range sentenceSplit[1:] {
		newSentence = w + " " + newSentence
	}
	return newSentence
}
