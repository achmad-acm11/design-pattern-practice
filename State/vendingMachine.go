package State

import "fmt"

type VendingMachine struct {
	HasItem       StateInterface
	ItemRequested StateInterface
	HasMoney      StateInterface
	NoItem        StateInterface

	CurrentState StateInterface

	ItemCount int
	ItemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}

	hasItemState := &HasItemState{VendingMachine: v}
	itemRequestedState := &ItemRequestedState{VendingMachine: v}
	hasMoneyState := &HasMoneyState{VendingMachine: v}
	noItemState := &NoItemState{VendingMachine: v}

	v.SetState(hasItemState)
	v.HasItem = hasItemState
	v.ItemRequested = itemRequestedState
	v.HasMoney = hasMoneyState
	v.NoItem = noItemState

	return v
}

func (m *VendingMachine) AddItem(i int) error {
	return m.CurrentState.AddItem(i)
}

func (m *VendingMachine) RequestItem() error {
	return m.CurrentState.RequestItem()
}

func (m *VendingMachine) InsertMoney(money int) error {
	return m.CurrentState.InsertMoney(money)
}

func (m *VendingMachine) DispenseItem() error {
	return m.CurrentState.DispenseItem()
}

func (m *VendingMachine) IncrementItemCount(count int) {
	fmt.Printf("Adding %d items \n", count)
	m.ItemCount = m.ItemCount + count
}

func (m *VendingMachine) SetState(s StateInterface) {
	m.CurrentState = s
}
