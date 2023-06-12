package main

import (
	"fmt"
	"strings"
)

type Set struct {
	s map[string]bool
}

func (set *Set) Add(value string) {
	set.s[value] = true // Добавляем заданное значение в множество, устанавливая значение true в карте.
}

func (set *Set) Print() {
	items := make([]string, 0, len(set.s)) // Создаем срез для хранения элементов множества.
	for item := range set.s {              // Итерируемся по ключам в карте.
		items = append(items, item) // Добавляем каждый ключ в срез.
	}
	fmt.Printf("set[%s]\n", strings.Join(items, " ")) // Выводим множество, объединяя элементы среза строк с помощью пробелов.
}

func main() {
	list := []string{"cat", "cat", "dog", "cat", "tree"} // Создаем срез строк.
	set := Set{s: make(map[string]bool)}                 // Создаем новое множество с пустой картой.
	for _, v := range list {                             // Итерируемся по элементам среза.
		set.Add(v) // Добавляем каждый элемент в множество.
	}
	fmt.Println("Before:")
	fmt.Println(list)
	fmt.Println("After:")
	set.Print() // Выводим элементы множества.
}
