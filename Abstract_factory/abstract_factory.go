package Abstract_factory

type pricy interface {
	Price() string
}

type Dimension struct {
	Length, Width, Height int
}

type Chair interface {
	pricy
	IsIOTEnabled() bool
	IsSoft() bool
}

type CoffeTable interface {
	pricy
	Size() Dimension
	IsFoldable() bool
}

type SoftStyle string

const (
	EuropeStyle   SoftStyle = "EuropeStyle"
	AmericanStyle SoftStyle = "AmericanStyle"
)

type Sofa interface {
	pricy
	Style() SoftStyle
}

type FurnitureFactory interface {
	CreateChair() Chair
	CreateCoffeTable() CoffeTable
	CreateSofa() Sofa
}
