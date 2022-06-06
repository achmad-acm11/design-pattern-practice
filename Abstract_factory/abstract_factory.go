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

// Factory Sport
type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

type ISportFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
