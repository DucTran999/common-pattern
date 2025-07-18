package singleton_test

import (
	"patterns/creational/singleton"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetInstance(t *testing.T) {
	t.Parallel()
	// Get the singleton instance
	s1 := singleton.GetInstance()

	// Get the same instance again
	s2 := singleton.GetInstance()

	assert.Equal(t, s1, s2)
}
