package builder

type House struct {
	Walls   string
	Doors   int
	Windows int
	Garage  bool
}

type houseBuilder struct {
	house House
}

// NewHouseBuilder initializes a new houseBuilder
func NewHouseBuilder() *houseBuilder {
	return &houseBuilder{
		house: House{},
	}
}

// SetWalls sets the type of walls
func (b *houseBuilder) SetWalls(walls string) *houseBuilder {
	b.house.Walls = walls
	return b
}

// SetDoors sets the number of doors
func (b *houseBuilder) SetDoors(doors int) *houseBuilder {
	if doors < 0 {
		doors = 0
	}

	b.house.Doors = doors
	return b
}

// SetWindows sets the number of windows
func (b *houseBuilder) SetWindows(windows int) *houseBuilder {
	if windows < 0 {
		windows = 0
	}

	b.house.Windows = windows
	return b
}

// SetGarage sets whether the house has a garage
func (b *houseBuilder) SetGarage(hasGarage bool) *houseBuilder {
	b.house.Garage = hasGarage
	return b
}

// Build returns the final House object
func (b *houseBuilder) Build() House {
	return b.house
}
