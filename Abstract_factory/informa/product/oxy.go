package product

import "factory-pattern/Abstract_factory"

type Oxy struct {
}

func (o Oxy) Price() string {
	return "199000"
}

func (o Oxy) Size() Abstract_factory.Dimension {
	return Abstract_factory.Dimension{
		Length: 60,
		Width:  34,
		Height: 26,
	}
}

func (o Oxy) IsFoldable() bool {
	return true
}
