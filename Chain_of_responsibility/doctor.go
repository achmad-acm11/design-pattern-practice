package Chain_of_responsibility

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(patent *Patient) {
	if patent.docterCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(patent)
		return
	}
	fmt.Println("Doctor checking patient")
	patent.docterCheckUpDone = true
	d.next.Execute(patent)
}

func (d *Doctor) SetNext(department Department) {
	d.next = department
}
