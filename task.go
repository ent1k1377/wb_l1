package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu  sync.Mutex
	Map map[int]string
}

func (s *SafeMap) Set(key int, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Map[key] = value
}

func main() {
	safeMap := SafeMap{Map: make(map[int]string)}

	var wg sync.WaitGroup
	var keys []int
	func(keys *[]int) {
		for i := 63; i < 125; i++ {
			*keys = append(*keys, i)
		}
	}(&keys)
	wg.Add(len(keys))

	for _, key := range keys {
		go func(k int) {
			defer wg.Done()
			safeMap.Set(k, string(k))
		}(key)
	}

	wg.Wait()

	// Вывод значений map
	for key, value := range safeMap.Map {
		fmt.Printf("%d: %s\n", key, value)
	}
}
