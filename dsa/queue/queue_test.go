package queue_test

import (
	"patterns/dsa/queue"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Queue(t *testing.T) {
	q := queue.NewQueue(10) // Create queue with capacity 10

	q.Enqueue(2)
	q.Enqueue(4)

	// Dequeue and check value
	val, err := q.Dequeue()
	require.NoError(t, err, "Dequeue should not return error")
	assert.Equal(t, 2, val, "Expected first dequeued value to be 2")

	// Dequeue and check value
	val, err = q.Dequeue()
	require.NoError(t, err, "Dequeue should not return error")
	assert.Equal(t, 4, val, "Expected first dequeued value to be 4")

	// Check queue length
	assert.Equal(t, 0, q.Len(), "Expected queue length to be 1 after one dequeue")

	// queue empty should return error
	val, err = q.Dequeue()
	assert.ErrorIs(t, err, queue.ErrQueueEmpty)
}
