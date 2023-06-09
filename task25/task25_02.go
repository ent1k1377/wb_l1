package main

import (
	"fmt"
	"time"
)

func Sleep02(d time.Duration) time.Time {
	ticker := time.Tick(d)
	for done := range ticker {
		return done
	}
	return time.Now()
}

func main() {
	start := time.Now()
	Sleep02(3 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("Slept for", elapsed.Seconds(), "seconds")
}
