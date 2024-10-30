package tests

import (
	"testing"

	"martindotexe/utils/ds/queue"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := queue.New()

	assert.Nil(t, q.Peek(), "Empty queue should be empty")
	assert.Zero(t, q.Len(), "Empty queue should be zero")
	assert.Nil(t, q.Dequeue(), "Dequeueing a queue should return nil when empty")

	q.Enqueue(1)
	assert.Equal(t, 1, q.Peek(), "Should return the value that is added")
	assert.Equal(t, 1, q.Len(), "Should be length of one")

	q.Enqueue(2)
	assert.Equal(t, 1, q.Peek(), "Should return the first value that is added")

	assert.Equal(t, 2, q.Len(), "Should increase length if enqueue")
	assert.Equal(t, 1, q.Dequeue(), "Should dequeue the first value")
	assert.Equal(t, 1, q.Len(), "Should decrease length when dequeueing")

	assert.Equal(t, 2, q.Peek())

	assert.Equal(t, 2, q.Dequeue())

	assert.Nil(t, q.Dequeue())
	assert.Nil(t, q.Dequeue())
	assert.Zero(t, q.Len())
}
