package main

import "fmt"

// 1. Объявляем структуру Human с полями Name, Age
// 2. Добавляем метод Speak структуры Human, ресивер передается по указателю,
// что позволяет не создавать копию объекта и соответственно не выделять лишнею память
// 3. Объявляем структуру Action в которой встраиваем структуру Human,
//что позволяет структуре Action пользоваться всеми методами структуры Human
// 4. Добавляем метод Program структуры Action (смысл такой же как и в методе Speak)
// 5. В функции main создаем экземпляр структуры Action заполняем в ней поля (Human) и вызываем два метода Speak и Program,
// вызов метода Speak работает потому, что мы ранее встроили Human в Action

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
