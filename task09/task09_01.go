package main

import "fmt"

func main() {
	read := make(chan int)
	write := make(chan int)

	var numbers []int
	func() {
		for i := 0; i < 20; i++ {
			numbers = append(numbers, i)
		}
	}()

	go producer(numbers, read)
	go doubler(read, write)
	consumer(write)
}

func producer(numbers []int, read chan int) {
	defer close(read)
	for _, num := range numbers {
		read <- num
	}
}

func doubler(read chan int, write chan int) {
	defer close(write)
	for ch := range read {
		double := ch * 2
		write <- double
	}
}

func consumer(write chan int) {
	for ch := range write {
		fmt.Println(ch)
	}
	fmt.Println("exit")
}
