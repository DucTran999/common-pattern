package bloom_test

import (
	"patterns/dsa/bloom-filter"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBloomFilter(t *testing.T) {
	tests := []struct {
		name        string
		numOfItems  uint64
		fpRate      float64
		expectedErr error
	}{
		{
			name:        "valid input",
			numOfItems:  100_000,
			fpRate:      0.01,
			expectedErr: nil,
		},
		{
			name:        "invalid number of items",
			numOfItems:  0,
			fpRate:      0.01,
			expectedErr: bloom.ErrInvalidNumberOfItems,
		},
		{
			name:        "invalid false positive rate",
			numOfItems:  10_000,
			fpRate:      0,
			expectedErr: bloom.ErrInvalidFalsePositiveRate,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := bloom.NewBloomFilter(tc.numOfItems, tc.fpRate)

			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}

func TestBloomFilter(t *testing.T) {
	bf, err := bloom.NewBloomFilter(1_000_000, 0.02)
	require.NoError(t, err)
	input := "daniel@gmail.com"
	unexpectedInput := "daisy@gmail.com"

	bf.Add([]byte(input))

	require.True(t, bf.MightContain([]byte(input)))
	require.False(t, bf.MightContain([]byte(unexpectedInput)))
}
