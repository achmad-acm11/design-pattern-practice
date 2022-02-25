package Composite

type Employee interface {
	GetSalary() int
	GetSubordinateSalary() int
}

type CEO struct {
	Subordinate []Employee
}

func (C CEO) GetSalary() int {
	return 10
}

func (C CEO) GetSubordinateSalary() int {
	result := 0
	for _, employee := range C.Subordinate {
		result += employee.GetSubordinateSalary()
	}
	return C.GetSalary() + result
}

type VP struct {
	Subordinate []Employee
}

func (V VP) GetSalary() int {
	return 5
}

func (V VP) GetSubordinateSalary() int {
	result := 0

	for _, employee := range V.Subordinate {
		result += employee.GetSalary()
	}
	return V.GetSalary() + result
}

type Staff struct {
}

func (s Staff) GetSalary() int {
	return 1
}

func (s Staff) GetSubordinateSalary() int {
	return 0
}





