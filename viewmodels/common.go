package viewmodels

type Head struct {
	Title string
}

func NewHead(title string) *Head {
	return &Head{
		Title: title,
	}
}

type Header struct {
	Active string
}

func NewHeader(active string) *Header {
	return &Header{
		Active: active,
	}
}

type Footer struct {
}

func NewFooter() *Footer {
	return &Footer{}
}
