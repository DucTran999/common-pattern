package bloom

import (
	"hash/fnv"
	"math"
)

type BloomFilter interface {
	Add(data []byte)

	MightContain(data []byte) bool
}

type bloomFilter struct {
	bitset []bool
	k      uint64 // number of hash function
	m      uint64 // size of bitset
}

// Create a Bloom filter
func NewBloomFilter(numberOfItems uint64, fpRate float64) *bloomFilter {
	m := estimateM(numberOfItems, fpRate)
	k := estimateK(numberOfItems, m)

	return &bloomFilter{
		bitset: make([]bool, m),
		k:      k,
		m:      m,
	}
}

// estimateM calculates the optimal number of bits (m) required in the Bloom filter's bit array,
// given:
//
//   - n: the expected number of items to store
//   - p: the desired false positive probability (0 < p < 1)
//
// Returns the number of bits rounded up to the nearest integer using math.Ceil.
func estimateM(n uint64, p float64) uint64 {
	// Compute numerator: n * ln(p)
	numerator := float64(n) * math.Log(p)

	// Compute denominator: (ln(2))^2
	denominator := math.Pow(math.Log(2), 2)

	// Apply formula and use math.Ceil to avoid undersizing
	m := -numerator / denominator

	// Convert to uint64
	return uint64(math.Ceil(m))
}

// estimateK calculates the optimal number of hash functions (k)
// for a Bloom filter given:
//
//   - n: expected number of elements
//   - m: size of the bit array (in bits)
func estimateK(n, m uint64) uint64 {
	if n == 0 || m == 0 {
		return 1 // Avoid division by zero or invalid config
	}
	return uint64(math.Round((float64(m) / float64(n)) * math.Ln2))
}

func (bf *bloomFilter) Add(data []byte) {
	// h1 and h2 must be different to reduce collision
	h1 := bf.hash1(data)
	h2 := bf.hash2(data)

	// Mark bit
	for i := uint64(0); i < bf.k; i++ {
		pos := (h1 + i*h2) % bf.m
		bf.bitset[pos] = true
	}
}

func (bf *bloomFilter) MightContain(data []byte) bool {
	h1 := bf.hash1(data)
	h2 := bf.hash2(data)

	for i := uint64(0); i < bf.k; i++ {
		pos := (h1 + i*h2) % bf.m
		if !bf.bitset[pos] {
			return false
		}
	}

	return true
}

func (bf *bloomFilter) hash1(data []byte) uint64 {
	h1 := fnv.New64a() // use hash fvn 1 a
	_, _ = h1.Write(data)
	return h1.Sum64()
}

func (bf *bloomFilter) hash2(data []byte) uint64 {
	h1 := fnv.New64() // use hash fvn 1
	_, _ = h1.Write(data)
	return h1.Sum64()
}
