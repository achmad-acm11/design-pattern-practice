package Chain_of_responsibility

type Department interface {
	Execute(patent *Patient)
	SetNext(department Department)
}
