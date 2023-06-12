package main

import (
	"fmt"
	"sync"
)

// NewConcurrencyMap создает и возвращает новый экземпляр ConcurrencyMap с инициализированным полем Map.
func NewConcurrencyMap[K comparable, V any]() *ConcurrencyMap[K, V] {
	return &ConcurrencyMap[K, V]{
		Map: make(map[K]V),
	}
}

// ConcurrencyMap представляет структуру данных для конкурентного доступа к отображению Map.
type ConcurrencyMap[K comparable, V any] struct {
	Map map[K]V
	sync.Mutex
}

// Set устанавливает значение value по ключу key в отображении Map.
func (cm *ConcurrencyMap[K, V]) Set(key K, value V) {
	cm.Lock()
	defer cm.Unlock()
	cm.Map[key] = value
}

func main() {
	var list []int

	// Анонимная функция, которая заполняет список list значениями от 65 до 122
	func(list *[]int) {
		for i := 65; i < 123; i++ {
			*list = append(*list, i)
		}
	}(&list)

	cm := NewConcurrencyMap[int, string]() // Создание нового экземпляра ConcurrencyMap с ключами int и значениями string

	var wg sync.WaitGroup
	wg.Add(len(list))

	// Запуск горутин для каждого значения из списка list
	for _, v := range list {
		go func(v int) {
			defer wg.Done()

			cm.Set(v, string(rune(v)))                // Установка значения в ConcurrencyMap
			fmt.Printf("%d:%s\n", v, string(rune(v))) // Вывод значения
		}(v)
	}

	wg.Wait() // Ожидание завершения всех горутин

	// Проверка, что количество элементов в списке list равно количеству элементов в отображении Map в ConcurrencyMap
	if len(list) == len(cm.Map) {
		fmt.Println("\nit's all good")
	}
}
