package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func (a *Action) Program() {
	fmt.Println("wrk -t12 -c400 -d30s --latency http://localhost:3000/get-all-orders/")
}

func main() {
	action := Action{
		Human: Human{
			Name: "Daniel",
			Age:  21,
		},
	}

	action.Speak()
	action.Program()
}
