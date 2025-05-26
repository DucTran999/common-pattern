package hashtable

import (
	"log"
)

type HashTable interface {
	Find(val string) bool
	Insert(val string)
	Delete(val string)
}

// hash table structure
type hashTable struct {
	indexArray []*bucket
}

// SimpleHash converts a string to a hash code using ASCII values and modulus
func simpleHash(s string, buckets int) int {
	hash := 0
	for _, char := range s {
		hash += int(char) // Add ASCII value (or rune value for Unicode)
	}
	return hash % buckets // Modulus to fit bucket range
}

func (h *hashTable) Insert(val string) {
	idx := simpleHash(val, len(h.indexArray))
	h.indexArray[idx].insert(val)
}

func (h *hashTable) Delete(val string) {
	idx := simpleHash(val, len(h.indexArray))
	h.indexArray[idx].delete(val)
}

func (h *hashTable) Find(val string) bool {
	idx := simpleHash(val, len(h.indexArray))
	return h.indexArray[idx].find(val)
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(k string) {
	if !b.find(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		log.Println("value already existed")
	}
}

func (b *bucket) find(k string) bool {
	currNode := b.head

	for currNode != nil {
		if currNode.key == k {
			return true
		}
		currNode = currNode.next
	}

	return false
}

func (b *bucket) delete(k string) {
	// Handle empty bucket
	if b.head == nil {
		return
	}

	// If element need to find is root of bucket
	// unlink it by point to the next bucket node
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	// iterate to the leaf node of bucket
	for prevNode.next != nil {
		// If next node match value then ignore it and point to the next.next node
		if prevNode.next.key == k {
			// remove
			prevNode.next = prevNode.next.next
			return // Exit after deletion to avoid no used loop
		}

		prevNode = prevNode.next
	}
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init(arrayLen int) HashTable {
	if arrayLen == 0 {
		arrayLen = 10
	}

	ht := &hashTable{
		indexArray: make([]*bucket, arrayLen),
	}

	// Assign bucket
	for i := range ht.indexArray {
		ht.indexArray[i] = &bucket{}
	}

	return ht
}
