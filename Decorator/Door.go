package Decorator

import "fmt"

type Openner interface {
	Open()
}

type Door struct {
}

func (d Door) Open() {
	fmt.Println("Open the Door")
}

func NewDoorIOT(openner Openner) *DoorByIOT {
	return &DoorByIOT{
		openner: openner,
	}
}

type DoorByIOT struct {
	openner Openner
}

func (d DoorByIOT) OpenWithIOT() {
	fmt.Println("Open with technology")
	d.openner.Open()
}
