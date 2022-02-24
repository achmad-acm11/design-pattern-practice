package Prototype

type Prototype interface {
	Clone() Prototype
}

type Robot struct {
	Id         string
	Name       string
	processeor string
}

func (r Robot) Clone() Prototype {
	return &Robot{
		Id:         r.Id,
		Name:       r.Name,
		processeor: r.processeor,
	}
}
