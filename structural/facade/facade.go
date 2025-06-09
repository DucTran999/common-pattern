package facade

type television struct {
	dvd      ElectronicDevice
	proj     ElectronicDevice
	screen   ElectronicDevice
	soundSys ElectronicDevice
}

func NewTelevision() *television {
	return &television{
		dvd:      NewDVDPlayer(),
		proj:     NewProjector(),
		screen:   NewScreen(),
		soundSys: NewSoundSystem(),
	}
}

func (t *television) TurnOn() {
	t.dvd.TurnOn()
	t.soundSys.TurnOn()
	t.proj.TurnOn()
	t.screen.TurnOn()
}

func (t *television) TurnOff() {
	t.screen.TurnOff()
	t.proj.TurnOff()
	t.soundSys.TurnOff()
	t.dvd.TurnOff()
}

func (t *television) CheckAllDeviceOn() bool {
	return t.dvd.Status() && t.proj.Status() && t.screen.Status() && t.soundSys.Status()
}

func (t *television) CheckAllDeviceOff() bool {
	return !t.dvd.Status() && !t.proj.Status() && !t.screen.Status() && !t.soundSys.Status()
}
