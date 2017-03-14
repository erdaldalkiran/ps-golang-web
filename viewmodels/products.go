package viewmodels

type Products struct {
	Head    *Head
	Header  *Header
	Footer  *Footer
	Content interface{}
}

func NewProducts() *Products {
	return &Products{
		Head:   NewHead("Lemonade Stand Society"),
		Header: NewHeader("products"),
		Footer: NewFooter(),
	}
}
