package State

import "fmt"

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (h HasItemState) AddItem(i int) error {
	fmt.Printf("%d items added \n", i)
	h.VendingMachine.IncrementItemCount(i)
	return nil
}

func (h HasItemState) RequestItem() error {
	if h.VendingMachine.ItemCount == 0 {
		h.VendingMachine.SetState(h.VendingMachine.NoItem)
		return fmt.Errorf("No item present")
	}
	fmt.Println("Item requested")
	h.VendingMachine.SetState(h.VendingMachine.ItemRequested)
	return nil
}

func (h HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (h HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
