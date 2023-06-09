package main

import (
	"io"
	"log"
	"strings"
)

var justString []byte

func main() {
	someFunc()
}

func someFunc() {
	builder := strings.Builder{}
	v := createHugeString(&builder, 1<<10)
	justString = make([]byte, 100)
	copy(justString, v)
}

func createHugeString(builder *strings.Builder, length int) string {
	for i := 0; i < length; i++ {
		_, err := io.WriteString(builder, "a")
		if err != nil {
			log.Fatal(err)
		}
	}
	return builder.String()
}
