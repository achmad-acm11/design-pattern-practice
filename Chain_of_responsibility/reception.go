package Chain_of_responsibility

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Execute(patent *Patient) {
	if patent.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(patent)
		return
	}
	fmt.Println("Reception registering patient")
	patent.registrationDone = true
	r.next.Execute(patent)
}

func (r *Reception) SetNext(department Department) {
	r.next = department
}
