package viewmodels

type ProfileVM struct {
	Head    *Head
	Header  *Header
	Footer  *Footer
	Content *User
}

type User struct {
	Id        int
	Email     string
	FirstName string
	LastName  string
	Stand     Stand
}

type Stand struct {
	Address string
	City    string
	State   string
	Zip     string
}

func NewProfileVM() *ProfileVM {
	return &ProfileVM{
		Head:    NewHead("Lemonade Stand Society - Profile"),
		Header:  NewHeader("profile"),
		Footer:  NewFooter(),
		Content: &User{},
	}
}
