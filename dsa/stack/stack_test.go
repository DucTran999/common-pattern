package stack_test

import (
	"patterns/dsa/stack"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PopGotErrWhenStackEmpty(t *testing.T) {
	s := stack.NewStack()

	_, err := s.Pop()

	require.Error(t, err)
}

func Test_Push(t *testing.T) {
	expectedVal := 6
	expectedLen := 1

	s := stack.NewStack()

	s.Push(5)
	s.Push(6)

	val, err := s.Pop()

	require.NoError(t, err)
	require.Equal(t, expectedVal, val)
	assert.Equal(t, expectedLen, s.Len())
}
