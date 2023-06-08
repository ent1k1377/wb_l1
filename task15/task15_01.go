package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

var justString string

func someFunc() {
	length := 100
	var builder strings.Builder
	createHugeString(&builder, length)
	justString = builder.String()
}

func main() {
	someFunc()
	log.Println(justString)
}

func createHugeString(builder *strings.Builder, length int) {
	for i := 0; i < length; i++ {
		_, err := io.WriteString(builder, "a")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(builder.Len(), builder.Cap())
	}
}
