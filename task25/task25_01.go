package main

import (
	"fmt"
	"time"
)

func Sleep01(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	start := time.Now()
	Sleep01(3 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("Slept for", elapsed.Seconds(), "seconds")
}
