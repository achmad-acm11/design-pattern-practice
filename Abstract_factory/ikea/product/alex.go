package product

type Alex struct {
}

func (a Alex) Price() string {
	return "1096000"
}

func (a Alex) IsIOTEnabled() bool {
	return false
}

func (a Alex) IsSoft() bool {
	return false
}
