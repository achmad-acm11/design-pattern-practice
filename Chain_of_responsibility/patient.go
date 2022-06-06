package Chain_of_responsibility

type Patient struct {
	Name              string
	registrationDone  bool
	docterCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}
