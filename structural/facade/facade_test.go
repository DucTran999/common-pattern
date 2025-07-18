package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Television(t *testing.T) {
	t.Parallel()
	tv := NewTelevision()

	tv.TurnOn()

	require.True(t, tv.CheckAllDeviceOn(), "all device on should return true")

	tv.TurnOff()

	assert.True(t, tv.CheckAllDeviceOff(), "all device off should return true")
}
