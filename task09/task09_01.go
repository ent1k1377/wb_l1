package main

import "fmt"

func main() {
	read := make(chan int)  // Канал для чтения чисел из производителя
	write := make(chan int) // Канал для записи удвоенных чисел для потребителя

	var numbers []int
	func() {
		for i := 0; i < 20; i++ {
			numbers = append(numbers, i) // Создание среза чисел от 0 до 19
		}
	}()

	go producer(&numbers, read) // Запуск горутины производителя для отправки чисел в канал
	go doubler(read, write)     // Запуск горутины удвоителя для обработки чисел из канала
	consumer(write)             // Запуск потребителя для чтения и вывода удвоенных чисел
}

func producer(numbers *[]int, read chan int) {
	defer close(read) // Закрытие канала чтения после отправки всех чисел
	for _, num := range *numbers {
		read <- num // Отправка чисел в канал чтения
	}
}

func doubler(read chan int, write chan int) {
	defer close(write) // Закрытие канала записи после обработки всех чисел
	for r := range read {
		double := r * 2 // Удвоение числа
		write <- double // Отправка удвоенного числа в канал записи
	}
}

func consumer(write chan int) {
	for w := range write {
		fmt.Println(w) // Вывод удвоенных чисел
	}
	fmt.Println("exit") // Вывод сообщения о завершении работы
}
