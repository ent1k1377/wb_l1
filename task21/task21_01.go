package main

import "fmt"

type Client struct {
}

func (c Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts lightnig connector into computer")
	com.InsertIntoLightningPort()
}

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightnig connector is plugged into mac machine")
}

type Windows struct {
}

func (w Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts lightning signal to USB")
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
