package main

import "fmt"

func main() {
	// Здесь происходит то же самое, что и в 01 примере, но WaitGroup заменен на не буферизованный канал,
	// который записывает во внутренний буфер в каждой горутине пустую структуру (т.к. пустая структура(0 байт), по сравнению с bool(1 байт))
	numbers := []int{2, 4, 6, 8, 10}
	done := make(chan struct{})

	for _, n := range numbers {
		go func(num int) {
			square := num * num
			fmt.Printf("%d * %d = %d\n", num, num, square)
			done <- struct{}{}
		}(n)
	}

	// и в конце программы у нас блокируется main из-за канала, пока for range не проитерируется и канал не прочитает из буфера данные
	// все сломается если numbers увеличит свою длину, так что, это не безопасный способ
	for range numbers {
		<-done
	}

	fmt.Println("exit")
}
