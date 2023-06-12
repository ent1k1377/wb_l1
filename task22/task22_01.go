package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем объекты типа big. Int для хранения больших чисел
	a := big.NewInt(0)
	b := big.NewInt(0)

	// Устанавливаем значения переменных a и b
	a.SetString("1000000000123455129991000000", 10)
	b.SetString("9999999993214419", 10)

	fmt.Print(a, "\n")
	fmt.Print(b, "\n")

	// Умножение
	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление
	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("Деление:", div)

	// Сложение
	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Println("Сложение:", sum)

	// Вычитание
	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Println("Вычитание:", sub)
}
