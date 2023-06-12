package main

import (
	"fmt"
)

type MyStruct struct {
	Name string
	Age  int
}

func main() {
	numbers := func() []int {
		temp := make([]int, 100)
		for i := 0; i < len(temp); i++ {
			temp[i] = i
		}
		return temp
	}()
	done := make(chan struct{})

	for _, n := range numbers {
		go func(num int) {
			square := num * num
			fmt.Printf("%d * %d = %d\n", num, num, square)
			done <- struct{}{}
		}(n)
	}

	for range numbers {
		<-done
	}

	fmt.Println("exit")
}
