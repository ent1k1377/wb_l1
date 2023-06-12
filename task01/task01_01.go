package main

import "fmt"

// Human Объявляем структуру с полями Name, Age
type Human struct {
	Name string
	Age  int
}

// Speak Добавляем метод структуры Human, ресивер передается по указателю,
// что позволяет не создавать копию объекта и соответственно не выделять лишнею память
func (h *Human) Speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old\n", h.Name, h.Age)
}

// Action Объявляем структуру в которой встраиваем структуру Human,
// что позволяет структуре Action пользоваться всеми методами структуры Human
type Action struct {
	Human
}

// Program Добавляем метод структуры Action
func (a *Action) Program() {
	fmt.Println("wrk -t12 -c400 -d30s --latency http://localhost:3000/get-all-orders/")
}

// В функции main создаем экземпляр структуры Action заполняем в ней поля (Human) и вызываем два метода Speak и Program,
// вызов метода Speak работает потому, что мы ранее встроили Human в Action
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
