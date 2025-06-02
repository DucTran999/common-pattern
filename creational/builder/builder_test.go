package builder_test

import (
	"patterns/creational/builder"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_HouseBuilder(t *testing.T) {
	t.Run("FullyConfiguredHouse", func(t *testing.T) {
		house := builder.NewHouseBuilder().
			SetWindows(2).
			SetWalls("brick").
			SetDoors(1).
			SetGarage(true).
			Build()

		require.Equal(t, "brick", house.Walls, "Walls should be 'brick'")
		require.Equal(t, 2, house.Windows, "Windows should be 2")
		require.Equal(t, 1, house.Doors, "Doors should be 1")
		require.True(t, house.Garage, "Garage should be true")
	})

	t.Run("MinimalHouse", func(t *testing.T) {
		house := builder.NewHouseBuilder().
			SetWalls("wood").
			Build()

		require.Equal(t, "wood", house.Walls, "Walls should be 'wood'")
		require.Equal(t, 0, house.Windows, "Windows should be 0 (default)")
		require.Equal(t, 0, house.Doors, "Doors should be 0 (default)")
		require.False(t, house.Garage, "Garage should be false (default)")
	})

	t.Run("DefaultHouse", func(t *testing.T) {
		house := builder.NewHouseBuilder().
			SetWalls("wood").
			SetDoors(-1).
			SetWindows(-1).
			Build()

		require.Equal(t, "wood", house.Walls, "Walls should be 'wood'")
		require.Equal(t, 0, house.Windows, "Windows should be 0 (default)")
		require.Equal(t, 0, house.Doors, "Doors should be 0 (default)")
		require.False(t, house.Garage, "Garage should be false (default)")
	})
}
