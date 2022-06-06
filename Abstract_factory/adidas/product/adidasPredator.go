package product

type AdidasPredator struct {
	Logo string
	Size int
}

func (a *AdidasPredator) SetLogo(logo string) {
	a.Logo = logo
}

func (a *AdidasPredator) SetSize(size int) {
	a.Size = size
}

func (a *AdidasPredator) GetLogo() string {
	return a.Logo
}

func (a *AdidasPredator) GetSize() int {
	return a.Size
}
