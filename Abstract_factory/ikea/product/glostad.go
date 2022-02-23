package product

import "factory-pattern/Abstract_factory"

type Glostad struct {
}

func (g Glostad) Price() string {
	return "1795000"
}

func (g Glostad) Style() Abstract_factory.SoftStyle {
	return Abstract_factory.EuropeStyle
}
