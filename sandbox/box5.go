package main

import (
	"fmt"
	"unsafe"
)

func main() {
	empty := struct{}{}
	nonEmpty0 := struct{ closed bool }{}
	nonEmpty1 := struct{ name string }{}
	nonEmpty2 := struct{ age int }{}
	nonEmpty3 := struct{ price float64 }{}
	nonEmpty4 := struct{ children []int }{}
	nonEmpty5 := struct {
		key string `size:"32"`
	}{}
	nonEmpty6 := struct {
		children []float64
		asd      int
	}{}

	fmt.Println(unsafe.Sizeof(empty))     // 0
	fmt.Println(unsafe.Sizeof(nonEmpty0)) // 1
	fmt.Println(unsafe.Sizeof(nonEmpty1)) // 16
	fmt.Println(unsafe.Sizeof(nonEmpty2)) // 8
	fmt.Println(unsafe.Sizeof(nonEmpty3)) // 8
	fmt.Println(unsafe.Sizeof(nonEmpty4)) // 24
	fmt.Println(unsafe.Sizeof(nonEmpty5)) // 16
	fmt.Println(unsafe.Sizeof(nonEmpty6)) // 32
}
