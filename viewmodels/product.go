package viewmodels

type Product struct {
	Head    *Head
	Header  *Header
	Footer  *Footer
	Content interface{}
}

func NewProduct() *Product {
	return &Product{
		Head:   NewHead("Lemonade Stand Society"),
		Header: NewHeader("product"),
		Footer: NewFooter(),
	}
}
