package facade

type soundSystem struct {
	electronic
}

func NewSoundSystem() *soundSystem {
	return &soundSystem{}
}
