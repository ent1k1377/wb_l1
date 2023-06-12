package main

import (
	"fmt"
)

// 1. Main. Здесь используется канал для синхронизации и передачи квадратов чисел,
// в main запускается одна горутина в которой for range перебирает числа, и передает в канал строку, после всех итераций канал закрывается.
// В конце программы происходит чтение из канала до тех пор, пока канал не закроется.
// Также в этом примере вывод происходит последовательно, но из-за этого страдает производительность по сравнению с другими примерами

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	squareCh := make(chan string)

	go func() {
		for _, n := range numbers {
			square := n * n
			squareCh <- fmt.Sprintf("%d * %d = %d", n, n, square)
		}
		close(squareCh)
	}()

	for ch := range squareCh {
		fmt.Println(ch)
	}
	fmt.Println("exit")
}
