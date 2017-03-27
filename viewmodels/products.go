package viewmodels

import "fmt"

type ProductsVM struct {
	Head    *Head
	Header  *Header
	Footer  *Footer
	Content []Product
}

func NewProducts() *ProductsVM {
	return &ProductsVM{
		Head:    NewHead("Lemonade Stand Society"),
		Header:  NewHeader("products"),
		Footer:  NewFooter(),
		Content: Products,
	}
}

func NewProductsByCategoryId(id int) (*ProductsVM, error) {
	if id != 1 {
		return nil, fmt.Errorf("products with category id %d does not exist", id)
	}

	return &ProductsVM{
		Head:    NewHead("Lemonade Stand Society"),
		Header:  NewHeader("products"),
		Footer:  NewFooter(),
		Content: Products,
	}, nil
}
