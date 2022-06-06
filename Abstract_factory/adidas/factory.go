package adidas

import (
	"factory-pattern/Abstract_factory"
	"factory-pattern/Abstract_factory/adidas/product"
)

type Adidas struct {
}

func (a Adidas) MakeShoe() Abstract_factory.IShoe {
	adidasPredator := &product.AdidasPredator{}
	adidasPredator.SetLogo("Adidas")
	adidasPredator.SetSize(30)
	return adidasPredator
}

func (a Adidas) MakeShirt() Abstract_factory.IShirt {
	adidasClimactCool := &product.AdidasClimactCool{}
	adidasClimactCool.SetLogo("Adidas")
	adidasClimactCool.SetSize(44)
	return adidasClimactCool
}
