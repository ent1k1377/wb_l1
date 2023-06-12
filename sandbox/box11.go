package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	for i := 0; i < 5; i++ {
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i, wg)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
