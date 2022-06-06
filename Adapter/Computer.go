package Adapter

import "fmt"

type Computer interface {
	InsertToLightningPort()
}

type Client struct {
}

func (c Client) InsertToLightningPort(com Computer) {
	fmt.Println("Client plugin Lightning Connector")
	com.InsertToLightningPort()
}

type Imac struct {
}

func (i Imac) InsertToLightningPort() {
	fmt.Println("IMac plugin Lightning Connector")
}

type Windows struct {
}

func (w Windows) InsertToUSBPort() {
	fmt.Println("Window plugin USB Port")
}

type WindowAdapter struct {
	WindowsObj *Windows
}

func (w WindowAdapter) InsertToLightningPort() {
	fmt.Println("Converting Lightning to port USB")
	w.WindowsObj.InsertToUSBPort()
}
