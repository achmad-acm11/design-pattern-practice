package Bridge

import "fmt"

type Shape interface {
	Draw()
	Area() float64
}

type Circle struct {
	Color  Color
	Radius int
}

type Square struct {
	Length int
}

func (c Circle) Draw() {
	fmt.Println("Draw Circle with color" + c.Color.Hex())
}

func (c Circle) Area() float64 {
	//TODO implement me
	panic("implement me")
}

func (s Square) Draw() {
	//TODO implement me
	panic("implement me")
}

func (s Square) Area() float64 {
	//TODO implement me
	panic("implement me")
}
