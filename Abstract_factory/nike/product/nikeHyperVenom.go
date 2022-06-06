package product

type NikeHyperVenom struct {
	Logo string
	Size int
}

func (n *NikeHyperVenom) SetLogo(logo string) {
	n.Logo = logo
}

func (n *NikeHyperVenom) SetSize(size int) {
	n.Size = size
}

func (n NikeHyperVenom) GetLogo() string {
	return n.Logo
}

func (n NikeHyperVenom) GetSize() int {
	return n.Size
}
