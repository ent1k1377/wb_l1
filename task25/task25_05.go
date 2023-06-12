package main

import (
	"fmt"
	"time"
)

// NewTicker посылает в канал каждую секунду тик, а мы в бесконечном цикле ждем когда тик станет больше endTime
func Sleep05(duration time.Duration) {
	ticker := time.NewTicker(1 * time.Second)
	endTime := time.Now().Add(duration)

	for tick := range ticker.C {
		fmt.Println("Tick:", tick)
		if tick.After(endTime) {
			ticker.Stop()
			break
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println("Before sleep")
	Sleep05(time.Second * 5)
	fmt.Println("After sleep")
	fmt.Println(time.Since(start))
}
