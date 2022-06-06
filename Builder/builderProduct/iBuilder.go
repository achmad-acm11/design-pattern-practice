package builderProduct

import "factory-pattern/Builder/product"

type IBuilder interface {
	SetWindowType()
	SetFloor()
	SetDoorType()
	GetResult() *product.House
}
