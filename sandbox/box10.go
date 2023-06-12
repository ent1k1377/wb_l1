package main

import "fmt"

func update(p *int) {
	b := 2
	fmt.Println(p, &p, &b)
	p = &b
	// *p = b
}

func main() {
	a := 1
	p := &a
	fmt.Println(p, &p)
	fmt.Println(*p)
	update(p)
	fmt.Println(p)

	fmt.Println(*p)
}
