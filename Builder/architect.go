package Builder

import (
	"factory-pattern/Builder/builderProduct"
	"factory-pattern/Builder/product"
)

type Architect struct {
	builder builderProduct.IBuilder
}

func (a Architect) NewArchitect(b builderProduct.IBuilder) *Architect {
	return &Architect{
		builder: b,
	}
}

func (a *Architect) SetBuilder(b builderProduct.IBuilder) {
	a.builder = b
}

func (a Architect) BuildHouse() *product.House {
	a.builder.SetFloor()
	a.builder.SetDoorType()
	a.builder.SetWindowType()
	return a.builder.GetResult()
}
