package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
	mu   sync.Mutex
}

func main() {
	intCh := make(chan int)
	strCh := make(chan string, 1)

	go func() {
		intCh <- 34
	}()
	strCh <- "olleh"

	checkType(12)
	checkType("тевирп")
	checkType(true)
	checkType(<-intCh)
	checkType(<-strCh)
	checkType(make(chan User))
	checkType(make(chan interface{}))
}

func checkType(undefined interface{}) {
	switch v := undefined.(type) {
	default:
		fmt.Printf("Value: %v, type: %T\n", v, v)
	}
}
