package facade

type screen struct {
	electronic
}

func NewScreen() *screen {
	return &screen{}
}
