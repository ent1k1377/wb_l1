package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// Проблема кода связана с созданием среза justString из большой строки v.
// Даже если мы ограничиваем длину среза до 100 символов, полный массив данных v все равно остается в памяти,
// что может привести к утечке памяти.
//
// Чтобы исправить это, мы используем функцию copy, чтобы создать новый срез justString с нужными данными.
// И мы не ссылаемся на оригинальный массив данных v, и лишние данные будут корректно освобождены.

func main() {
	// justString как срез байтов
	someFunc1()

	// justString как строка
	someFunc2()
}

func someFunc1() {
	var justString []byte

	// Создаем strings.Builder для построения огромной строки
	builder := strings.Builder{}
	v := createHugeString(&builder, 1<<10)

	// Создаем срез байтов размером 100 и копируем первые 100 байт из v в него
	justString = make([]byte, 100)
	copy(justString, v[0:100])
}

func someFunc2() {
	var justString string

	// Создаем огромную строку с помощью createHugeString2
	v := createHugeString2(1 << 10)

	// Создаем новый срез строк размером 100 и копируем первые 100 элементов из v в него
	newStr := make([]string, 100)
	copy(newStr, v)

	// Объединяем элементы среза строк в одну строку с помощью strings.Join
	justString = strings.Join(newStr, "")
	fmt.Println(justString)
}

func createHugeString(builder *strings.Builder, length int) string {
	// Строим огромную строку из символов 'a' с помощью strings.Builder
	for i := 0; i < length; i++ {
		_, err := io.WriteString(builder, "a")
		if err != nil {
			log.Fatal(err)
		}
	}
	return builder.String()
}

func createHugeString2(length int) []string {
	// Создаем срез строк размером length, где каждый элемент содержит строку "a"
	strSlice := make([]string, length)
	for i := 0; i < length; i++ {
		strSlice[i] = "a"
	}
	return strSlice
}
