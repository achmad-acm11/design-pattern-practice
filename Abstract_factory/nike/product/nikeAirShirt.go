package product

type NikeAirShirt struct {
	Logo string
	Size int
}

func (n *NikeAirShirt) SetLogo(logo string) {
	n.Logo = logo
}

func (n *NikeAirShirt) SetSize(size int) {
	n.Size = size
}

func (n NikeAirShirt) GetLogo() string {
	return n.Logo
}

func (n NikeAirShirt) GetSize() int {
	return n.Size
}
