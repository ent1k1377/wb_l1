package main

import (
	"fmt"
	"time"
)

// Tick - удобная оболочка для NewTicker, предоставляющая доступ только к тиковому каналу. Ждем завершения
func Sleep02(duration time.Duration) time.Time {
	ticker := time.Tick(duration)
	for done := range ticker {
		return done
	}
	return time.Now()
}

func main() {
	start := time.Now()
	fmt.Println("Before sleep")
	Sleep02(time.Second * 2)
	fmt.Println("After sleep")
	fmt.Println(time.Since(start))
}
