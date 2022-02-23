package ikea

import (
	"factory-pattern/Abstract_factory"
	"factory-pattern/Abstract_factory/ikea/product"
)

type Ikea struct {
}

func (i Ikea) CreateChair() Abstract_factory.Chair {
	return product.Alex{}
}

func (i Ikea) CreateCoffeTable() Abstract_factory.CoffeTable {
	return product.Lack{}
}

func (i Ikea) CreateSofa() Abstract_factory.Sofa {
	return product.Glostad{}
}
