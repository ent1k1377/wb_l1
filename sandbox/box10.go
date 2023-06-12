package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b
	fmt.Println(*p, &p, p, &b)
	*p = b
	fmt.Println(*p, &p, p, &b)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}
