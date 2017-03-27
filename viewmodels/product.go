package viewmodels

import "fmt"

var Products = []Product{
	Product{
		Name:             "Apple Juice",
		DescriptionShort: "The perfect blend of Washington apples.",
		DescriptionLong:  "The perfect blend of Washington apples.",
		PricePerLiter:    0.89,
		PricePer10Liter:  0.85,
		Origin:           "Ohio",
		IsOrganic:        true,
		ImageUrl:         "apple",
		Id:               2,
	},
	Product{
		Name:             "Watermelon Juice",
		DescriptionShort: "From sun-drenched fields in Florida.",
		DescriptionLong:  "From sun-drenched fields in Florida.",
		PricePerLiter:    3.99,
		PricePer10Liter:  3.84,
		Origin:           "Florida",
		IsOrganic:        true,
		ImageUrl:         "watermelon",
		Id:               3,
	},
	Product{
		Name:             "Kiwi Juice",
		DescriptionShort: "California sunshine and rain distilled into sweet goodness",
		DescriptionLong:  "California sunshine and rain distilled into sweet goodness",
		PricePerLiter:    5.99,
		PricePer10Liter:  5.54,
		Origin:           "California",
		IsOrganic:        false,
		ImageUrl:         "kiwi",
		Id:               4,
	},
	Product{
		Name:             "Mangosteen Juice",
		DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
		DescriptionLong:  "Our latest taste sensation, imported directly from Hawaii",
		PricePerLiter:    6.87,
		PricePer10Liter:  6.79,
		Origin:           "Hawaii",
		IsOrganic:        false,
		ImageUrl:         "mangosteen",
		Id:               5,
	},
	Product{
		Name:             "Orange Juice",
		DescriptionShort: "Fresh squeezed from Florida's best oranges.",
		DescriptionLong:  "Fresh squeezed from Florida's best oranges.",
		PricePerLiter:    1.67,
		PricePer10Liter:  1.63,
		Origin:           "Florida",
		IsOrganic:        false,
		ImageUrl:         "orange",
		Id:               6,
	},
	Product{
		Name:             "Pineapple Juice",
		DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
		DescriptionLong:  "An exotic and refreshing offering. Straight from Hawaii.",
		PricePerLiter:    2.48,
		PricePer10Liter:  2.27,
		Origin:           "Hawaii",
		IsOrganic:        false,
		ImageUrl:         "pineapple",
		Id:               7,
	},
	Product{
		Name:             "Strawberry Juice",
		DescriptionShort: "MThe perfect balance of sweet and tart.",
		DescriptionLong:  "The perfect balance of sweet and tart.",
		PricePerLiter:    4.36,
		PricePer10Liter:  4.27,
		Origin:           "California",
		IsOrganic:        false,
		ImageUrl:         "strawberry",
		Id:               8,
	},
}

type ProductVM struct {
	Head    *Head
	Header  *Header
	Footer  *Footer
	Content *Product
}

type Product struct {
	Id               int
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageUrl         string
}

func NewProductVM(id int) (*ProductVM, error) {
	p, err := getProductById(id)
	if err != nil {
		return nil, err
	}
	return &ProductVM{
		Head:    NewHead("Lemonade Stand Society"),
		Header:  NewHeader("product"),
		Footer:  NewFooter(),
		Content: p,
	}, nil
}

func getProductById(id int) (*Product, error) {
	for _, p := range Products {
		if p.Id == id {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("the product with id %d does not exist", id)
}
