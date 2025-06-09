package facade

type ElectronicDevice interface {
	TurnOn()
	TurnOff()
	Status() bool
}
