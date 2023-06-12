package main

import (
	"fmt"
	"time"
)

// NewTimer шлет один раз в канал сообщение о том что время прошло
func Sleep04(duration time.Duration) {
	timer := time.NewTimer(duration)
	for range timer.C {
		return
	}
}

func main() {
	start := time.Now()
	fmt.Println("Before sleep")
	Sleep04(time.Second * 2)
	fmt.Println("After sleep")
	fmt.Println(time.Since(start))
}
