package main

import (
	"fmt"
	"time"
)

// Здесь мы через бесконечный цикл ждем когда текущее время станет больше target
func Sleep03(duration time.Duration) {
	start := time.Now()
	target := start.Add(duration)

	for time.Now().Before(target) {
		// Ожидание до достижения указанного времени
	}
}

func main() {
	start := time.Now()
	fmt.Println("Before sleep")
	Sleep03(time.Second * 2)
	fmt.Println("After sleep")
	fmt.Println(time.Since(start))
}
