package hashtable_test

import (
	hashtable "patterns/dsa/hash-table"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Insert_Delete(t *testing.T) {
	ht := hashtable.Init(10) // Initialize hash table with 10 buckets

	// Insert values that should collide (same bucket via custom hash)
	keys := []string{"abcx", "cbax", "bacx", "xbac", "bxac", "baxc"}
	for _, key := range keys {
		ht.Insert(key)
		require.True(t, ht.Find(key), "Expected to find inserted key: %s", key)
	}

	// Delete non-existent key
	ht.Delete("xabx") // does not exist
	require.False(t, ht.Find("xabx"), "Expected false for non-existent key")

	// Insert duplicate key: should log "value already existed"
	ht.Insert("abcx") // depending on implementation, might print or ignore silently

	// Delete tail node of bucket chain
	ht.Delete("bacx")
	require.False(t, ht.Find("bacx"), "Expected not to find deleted tail key")

	// Delete middle node of bucket chain
	ht.Delete("cbax")
	require.False(t, ht.Find("cbax"), "Expected not to find deleted middle key")

	// Delete head node of bucket chain
	ht.Delete("abcx")
	require.False(t, ht.Find("abcx"), "Expected not to find deleted head key")

	// Remaining key should still be found
	assert.True(t, ht.Find("xbac"), "Remaining key should still be found")
}

func Test_DeleteAKeyManyTimes(t *testing.T) {
	ht := hashtable.Init(0)
	val1 := "val1"

	ht.Insert(val1)
	ht.Delete(val1)

	// Try delete again should be ignore
	ht.Delete(val1)
}
