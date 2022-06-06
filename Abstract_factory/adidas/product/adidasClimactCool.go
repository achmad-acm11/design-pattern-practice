package product

type AdidasClimactCool struct {
	Logo string
	Size int
}

func (a *AdidasClimactCool) SetLogo(logo string) {
	a.Logo = logo
}

func (a *AdidasClimactCool) SetSize(size int) {
	a.Size = size
}

func (a AdidasClimactCool) GetLogo() string {
	return a.Logo
}

func (a AdidasClimactCool) GetSize() int {
	return a.Size
}
