package hashtable_test

import (
	hashtable "patterns/dsa/hash-table"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Insert(t *testing.T) {
	ht := hashtable.Init(10)
	ht.Insert("daniel")

	// should stdout value already existed
	ht.Insert("daniel")

	require.True(t, ht.Find("daniel"))
}

func Test_Delete(t *testing.T) {
	val := "daniel"
	ht := hashtable.Init(10)

	ht.Insert(val)
	require.True(t, ht.Find(val), "must find value after inserted value")

	ht.Delete(val)
	assert.False(t, ht.Find(val), "must return false after deleted value")
}
