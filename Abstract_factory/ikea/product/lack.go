package product

import "factory-pattern/Abstract_factory"

type Lack struct {
}

func (l Lack) Price() string {
	return "799000"
}

func (l Lack) Size() Abstract_factory.Dimension {
	return Abstract_factory.Dimension{
		Length: 91,
		Width:  66,
		Height: 60,
	}
}

func (l Lack) IsFoldable() bool {
	return false
}
