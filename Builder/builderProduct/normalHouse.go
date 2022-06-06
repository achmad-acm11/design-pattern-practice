package builderProduct

import "factory-pattern/Builder/product"

type NormalHouse struct {
	windowType string
	doorType   string
	floor      int
}

func (h NormalHouse) NewNormalHouse() *NormalHouse {
	return &NormalHouse{}
}

func (h *NormalHouse) SetWindowType() {
	h.windowType = "Wooden Window"
}

func (h *NormalHouse) SetFloor() {
	h.floor = 2
}

func (h *NormalHouse) SetDoorType() {
	h.doorType = "Wooden Door"
}

func (h NormalHouse) GetResult() *product.House {
	return &product.House{
		WindowType: h.windowType,
		DoorType:   h.doorType,
		Floor:      h.floor,
	}
}
