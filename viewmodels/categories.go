package viewmodels

type Categories struct {
	Head    *Head
	Header  *Header
	Footer  *Footer
	Content []Category
}

type Category struct {
	Id            int
	ImageUrl      string
	Title         string
	Description   string
	IsOrientRight bool
}

func NewCategories() *Categories {
	return &Categories{
		Head:   NewHead("Lemonade Stand Society - Shop"),
		Header: NewHeader("shop"),
		Footer: NewFooter(),
		Content: []Category{
			Category{
				Id:       1,
				ImageUrl: "lemon",
				Title:    "Juices and Mixes",
				Description: `Explore our wide assortment of juices and mixes expected by
							today's lemonade stand clientelle. Now featuring a full line of
							organic juices that are guaranteed to be obtained from trees that
							have never been treated with pesticides or artificial
							fertilizers.`,
				IsOrientRight: false,
			},
			Category{
				Id:       2,
				ImageUrl: "kiwi",
				Title:    "Cups, Straws, and Other Supplies",
				Description: `From paper cups to bio-degradable plastic to straws and
						napkins, LSS is your source for the sundries that keep your stand
						running smoothly.`,
				IsOrientRight: true,
			},
			Category{
				Id:       3,
				ImageUrl: "pineapple",
				Title:    "Signs and Advertising",
				Description: `Sure, you could just wait for people to find your stand
						along the side of the road, but if you want to take it to the next
						level, our premium line of advertising supplies.`,
				IsOrientRight: false,
			},
		},
	}
}
