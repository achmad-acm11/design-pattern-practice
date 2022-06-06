package builderProduct

import "factory-pattern/Builder/product"

type IglooHouse struct {
	windowType string
	doorType   string
	floor      int
}

func (h IglooHouse) NewIglooHouse() *IglooHouse {
	return &IglooHouse{}
}

func (h *IglooHouse) SetWindowType() {
	h.windowType = "Snow Window"
}

func (h *IglooHouse) SetFloor() {
	h.floor = 1
}

func (h *IglooHouse) SetDoorType() {
	h.doorType = "Snow Door"
}

func (h IglooHouse) GetResult() *product.House {
	return &product.House{
		WindowType: h.windowType,
		DoorType:   h.doorType,
		Floor:      h.floor,
	}
}
