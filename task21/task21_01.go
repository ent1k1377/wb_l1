package main

import "fmt"

// Client представляет клиента, который вставляет разъем Lightning в компьютер.
type Client struct {
}

// InsertLightningConnectorIntoComputer вставляет разъем Lightning в компьютер.
func (c Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Клиент вставляет разъем Lightning в компьютер")
	com.InsertIntoLightningPort()
}

// Computer представляет интерфейс компьютера.
type Computer interface {
	InsertIntoLightningPort()
}

// Mac представляет компьютер Mac.
type Mac struct {
}

// InsertIntoLightningPort вставляет разъем Lightning в порт компьютера Mac.
func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Разъем Lightning вставлен в компьютер Mac")
}

// Windows представляет компьютер с операционной системой Windows.
type Windows struct {
}

// insertIntoUSBPort вставляет разъем USB в порт компьютера Windows.
func (w Windows) insertIntoUSBPort() {
	fmt.Println("Разъем USB вставлен в компьютер с Windows")
}

// WindowsAdapter представляет адаптер для компьютера с Windows.
type WindowsAdapter struct {
	windowsMachine *Windows
}

// InsertIntoLightningPort вставляет разъем Lightning в порт компьютера, используя адаптер для Windows.
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Адаптер преобразует сигнал Lightning в USB")
	w.windowsMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
	macMachine := &Mac{}

	client.InsertLightningConnectorIntoComputer(macMachine)

	windowsMachine := &Windows{}
	windowsAdapter := &WindowsAdapter{windowsMachine: windowsMachine}

	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
