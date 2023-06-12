package main

import (
	"fmt"
	"time"
)

func Sleep01(duration time.Duration) {
	// After ожидает истечения продолжительности duration, а затем отправляет текущее время по возвращаемому каналу.
	<-time.After(duration)
}

func main() {
	start := time.Now()
	fmt.Println("Before sleep")
	Sleep01(time.Second * 2)
	fmt.Println("After sleep")
	fmt.Println(time.Since(start))
}
