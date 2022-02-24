package Bridge

type Color interface {
	Hex() string
}

type Red struct {
}

type Blue struct {
}

func (r Red) Hex() string {
	return "red"
}

func (b Blue) Hex() string {
	return "blue"
}
