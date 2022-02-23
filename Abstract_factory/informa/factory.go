package informa

import (
	"factory-pattern/Abstract_factory"
	"factory-pattern/Abstract_factory/informa/product"
)

type Informa struct {
}

func (i Informa) CreateChair() Abstract_factory.Chair {
	return nil
}

func (i Informa) CreateCoffeTable() Abstract_factory.CoffeTable {
	return product.Oxy{}
}

func (i Informa) CreateSofa() Abstract_factory.Sofa {
	return nil
}
