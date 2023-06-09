package main

import (
	"fmt"
	"strings"
)

type Set struct {
	s map[string]bool
}

func (s *Set) Print() {
	var output []string
	var i int
	for k, _ := range s.s {
		output = append(output, k)
		i++
	}
	fmt.Println("set[" + strings.Join(output, " ") + "]")
}

func (s *Set) Add(value string) {
	s.s[value] = true
}

func (s *Set) Intersection(q *Set) {
	for k, _ := range s.s {
		if value, _ := q.s[k]; !value {
			delete(s.s, k)
		}
	}
}

func main() {
	set1 := Set{s: make(map[string]bool)}
	set2 := Set{s: make(map[string]bool)}

	list1 := []string{"go", "python", "c", "go"}
	list2 := []string{"go", "python", "c++", "python"}

	for _, v := range list1 {
		set1.Add(v)
	}
	for _, v := range list2 {
		set2.Add(v)
	}

	fmt.Println("set1 and set2:")
	set1.Print()
	set2.Print()

	set1.Intersection(&set2)

	fmt.Println("\nset1: After intersection:")
	set1.Print()
}
