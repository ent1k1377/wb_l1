package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func worker(id int, message <-chan string) {
	// Функция воркера, которая получает сообщения из канала и обрабатывает их
	for m := range message {
		fmt.Printf("id: %d, message: %s\n", id, m)
	}
}

func main() {
	fmt.Print("Введите число воркеров: ")
	var numberWorker int
	_, err := fmt.Scan(&numberWorker)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	message := make(chan string, numberWorker) // Канал для передачи сообщений между главной горутиной и воркерами

	for i := 0; i < numberWorker; i++ {
		go worker(i, message) // Запуск воркеров
	}

	// Горутина для чтения ввода пользователя и отправки сообщений в канал
	go func() {
		for {
			var input string
			fmt.Scan(&input) // Чтение ввода пользователя
			message <- input // Отправка сообщения в канал
		}
	}()

	stop := make(chan os.Signal, 1)
	go signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM) // Регистрация обработчика сигналов SIGINT и SIGTERM

	<-stop // Ожидание получения сигнала

	close(message)        // Закрытие канала
	fmt.Println("\nexit") // Вывод сообщения о завершении программы
}
