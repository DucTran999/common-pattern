package facade

type electronic struct {
	isOn bool
}

func (s *electronic) TurnOn() {
	s.isOn = true
}

func (s *electronic) TurnOff() {
	s.isOn = false
}

func (s *electronic) Status() bool {
	return s.isOn
}
