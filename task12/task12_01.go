package main

import (
	"fmt"
	"strings"
)

type Set struct {
	s map[string]bool
}

func (s *Set) Add(value string) {
	s.s[value] = true
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

func main() {
	list := []string{"cat", "cat", "dog", "cat", "tree"}
	set := Set{s: make(map[string]bool)}
	for _, v := range list {
		set.Add(v)
	}

	set.Print()
}
