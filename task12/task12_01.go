package main

import (
	"fmt"
)

type Set struct {
	s map[string]bool
}

func (s *Set) Print() {
	q := "set["
	for k, _ := range s.s {
		q += k + " "
	}
	fmt.Println(q + "]")
}

func main() {
	list := []string{"cat", "cat", "dog", "cat", "tree"}
	set := Set{s: make(map[string]bool)}
	for _, v := range list {
		set.s[v] = true
	}

	set.Print()
}
