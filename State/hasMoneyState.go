package State

import "fmt"

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (h HasMoneyState) AddItem(i int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (h HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (h HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (h HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing item")
	h.VendingMachine.ItemCount = h.VendingMachine.ItemCount - 1
	if h.VendingMachine.ItemCount == 0 {
		h.VendingMachine.SetState(h.VendingMachine.NoItem)
	} else {
		h.VendingMachine.SetState(h.VendingMachine.HasItem)
	}
	return nil
}
