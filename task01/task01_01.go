package main

import "fmt"

// Human Определение структуры Human
type Human struct {
	Name string
	Age  int
}

// Speak Метод структуры Human
func (h *Human) Speak() {
	fmt.Println("I am a human.")
}

// Action Определение структуры Action
type Action struct {
	Human // Встраивание структуры Human в структуру Action
}

// Run Метод структуры Action
func (a *Action) Run() {
	fmt.Println("I can run.")
}

func main() {
	// Создание объекта структуры Action
	action := Action{
		Human: Human{
			Name: "John",
			Age:  25,
		},
	}

	// Доступ к методам из структуры Human через структуру Action
	action.Speak() // Выводит: "I am a human."
	action.Run()   // Выводит: "I can run."
}
