package main

import (
	"fmt"
	"strings"
)

type Set struct {
	s map[string]struct{} // s - это карта (map), которая хранит строки в качестве ключей с пустой структурой в качестве значения.
}

func NewSet() *Set {
	return &Set{
		s: make(map[string]struct{}), // Создаем и инициализируем новое множество с пустой картой.
	}
}

func (set *Set) Print() {
	items := make([]string, 0, len(set.s)) // Создаем срез для хранения элементов множества.
	for item := range set.s {              // Итерируемся по ключам в карте.
		items = append(items, item) // Добавляем каждый ключ в срез.
	}
	fmt.Printf("set[%s]\n", strings.Join(items, " ")) // Выводим множество, объединяя элементы среза строк с помощью пробелов.
}

func (set *Set) Add(value string) {
	set.s[value] = struct{}{} // Добавляем заданное значение в множество, присваивая пустую структуру в качестве значения в карте.
}

func (set *Set) Intersection(q *Set) {
	for k := range set.s { // Итерируемся по ключам в множестве.
		if !q.Contains(k) { // Проверяем, существует ли ключ в другом множестве.
			delete(set.s, k) // Если нет, удаляем ключ из множества.
		}
	}
}

func (set *Set) Contains(key string) bool {
	_, exists := set.s[key] // Проверяем, существует ли заданный ключ в множестве.
	return exists
}

func main() {
	set1 := NewSet() // Создаем новое множество.
	set2 := NewSet()

	list1 := []string{"go", "python", "c", "go"} // Создаем срез строк.
	list2 := []string{"go", "python", "c++", "python"}

	for _, v := range list1 { // Итерируемся по элементам среза.
		set1.Add(v) // Добавляем каждый элемент в set1.
	}
	for _, v := range list2 {
		set2.Add(v)
	}

	fmt.Println("set1 and set2:")
	set1.Print() // Выводим элементы множества set1.
	set2.Print() // Выводим элементы множества set2.

	set1.Intersection(set2) // Выполняем операцию пересечения между set1 и set2.

	fmt.Println("\nset1: After intersection:")
	set1.Print() // Выводим элементы множества set1 после операции пересечения.
}
