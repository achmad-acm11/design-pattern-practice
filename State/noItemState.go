package State

import "fmt"

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (n NoItemState) AddItem(i int) error {
	n.VendingMachine.IncrementItemCount(i)
	n.VendingMachine.SetState(n.VendingMachine.HasItem)
	return nil
}

func (n NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}

func (n NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (n NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}
