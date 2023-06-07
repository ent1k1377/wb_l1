package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Before sleep")
	sleep(2 * time.Second)
	fmt.Println("After sleep")
}
