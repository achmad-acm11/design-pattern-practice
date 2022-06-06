package nike

import (
	"factory-pattern/Abstract_factory"
	"factory-pattern/Abstract_factory/nike/product"
)

type Nike struct {
}

func (n Nike) MakeShoe() Abstract_factory.IShoe {
	hyperVenom := &product.NikeHyperVenom{}
	hyperVenom.SetSize(44)
	hyperVenom.SetLogo("Nike")
	return hyperVenom
}

func (n Nike) MakeShirt() Abstract_factory.IShirt {
	return &product.NikeAirShirt{
		Size: 30,
		Logo: "Nike",
	}
}
