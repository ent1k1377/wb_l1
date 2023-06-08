package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	timer := time.NewTimer(duration)

	select {
	case <-timer.C:
		return
	case <-time.After(duration):
		return
	}
}

func main() {
	start := time.Now()
	Sleep(3 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("Slept for", elapsed.Seconds(), "seconds")
}
