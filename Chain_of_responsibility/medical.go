package Chain_of_responsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(patent *Patient) {
	if patent.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(patent)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	patent.medicineDone = true
	m.next.Execute(patent)
}

func (m *Medical) SetNext(department Department) {
	m.next = department
}
