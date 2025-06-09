package facade

type projector struct {
	electronic
}

func NewProjector() *projector {
	return &projector{}
}
