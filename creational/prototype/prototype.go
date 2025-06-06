package prototype

type Document struct {
	Title string
	Body  string
}

func (d *Document) Clone() *Document {
	clone := *d
	return &clone
}
