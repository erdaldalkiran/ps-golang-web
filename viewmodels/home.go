package viewmodels

type Home struct {
	Head    *Head
	Header  *Header
	Footer  *Footer
	Content interface{}
}

func NewHome() *Home {
	return &Home{
		Head:   NewHead("Lemonade Stand Society"),
		Header: NewHeader("home"),
		Footer: NewFooter(),
	}
}
