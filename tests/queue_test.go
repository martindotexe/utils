package tests

import (
	"fmt"
	"testing"

	"github.com/martindotexe/utils/ds/queue"
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

func TestIterQueue(t *testing.T) {
	s1 := queue.New()

	s1.Enqueue("1")
	s1.Enqueue("2")
	s1.Enqueue("3")

	for i, n := range s1.Items() {
		assert.Equal(t, fmt.Sprintf("%d", i+1), n)
	}

	s2 := queue.New()

	for _, n := range s2.Items() {
		assert.Nil(t, n)
		assert.Zero(t, n)
		fmt.Print("Should not print")
	}
}
