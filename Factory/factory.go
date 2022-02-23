package Factory

import "fmt"

type Type interface {
	Running()
}

type Street struct {
}
type Water struct {
}

func (s *Street) Running() {
	fmt.Println("Transportation is running on street")
}

func (w *Water) Running() {
	fmt.Println("Transportation is running on water")
}

type Product interface {
	DoStuff() Type
}

type Car struct {
}

type Boat struct {
}

func (c Car) DoStuff() Type {
	return &Street{}
}

func (b Boat) DoStuff() Type {
	return &Water{}
}
