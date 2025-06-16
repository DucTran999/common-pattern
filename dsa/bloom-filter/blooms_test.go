package bloom_test

import (
	"patterns/dsa/bloom-filter"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBloomFilter(t *testing.T) {
	bf := bloom.NewBloomFilter(1_000_000, 0.02)
	input := "daniel@gmail.com"
	unexpectedInput := "daisy@gmail.com"

	bf.Add([]byte(input))

	require.True(t, bf.MightContain([]byte(input)))
	require.False(t, bf.MightContain([]byte(unexpectedInput)))
}
