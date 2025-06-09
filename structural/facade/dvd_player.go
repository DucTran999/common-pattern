package facade

type dvdPlayer struct {
	electronic
}

func NewDVDPlayer() *dvdPlayer {
	return &dvdPlayer{}
}
