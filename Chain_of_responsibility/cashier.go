package Chain_of_responsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(patent *Patient) {
	if patent.paymentDone {
		fmt.Println("Patient payment done")
	}
	fmt.Println("Cashier getting money form patient")
	patent.paymentDone = true
}

func (c *Cashier) SetNext(department Department) {
	c.next = department
}
