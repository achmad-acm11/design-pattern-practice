package State

import "fmt"

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i ItemRequestedState) AddItem(i2 int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (i ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item already requested")
}

func (i ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.ItemPrice {
		return fmt.Errorf("Inserted money is less. Please insert %d", i.VendingMachine.ItemPrice)
	}
	fmt.Println("Money entered is ok")
	i.VendingMachine.SetState(i.VendingMachine.HasMoney)
	return nil
}

func (i ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
