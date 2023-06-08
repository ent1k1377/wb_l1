package main

import (
	"fmt"
	"sync"
)

type ConcurrencyMap struct {
	Map map[any]any
	mu  sync.Mutex
}

func main() {
	var list []int
	func(list *[]int) {
		for i := 65; i < 123; i++ {
			*list = append(*list, i)
		}
	}(&list)

	m := ConcurrencyMap{
		Map: make(map[any]any),
	}
	var wg sync.WaitGroup
	wg.Add(len(list))
	for _, v := range list {
		go func(v int) {
			defer wg.Done()

			m.mu.Lock()
			m.Map[v] = string(rune(v))
			fmt.Printf("%d:%s\n", v, string(rune(v)))
			m.mu.Unlock()
		}(v)
	}

	wg.Wait()
	if len(list) == len(m.Map) {
		fmt.Println("\nit's all good")
	}
}
