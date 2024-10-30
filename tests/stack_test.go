package tests

import (
	"fmt"
	"testing"

	"github.com/martindotexe/utils/ds/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := stack.New()

	assert.Zero(t, s.Len(), "Stack length should be zero after initialized")

	assert.True(t, s.IsEmpty(), "Stack should be empty after initialized")

	assert.Nil(t, s.Pop(), "Stack should return nil value if empty")

	assert.Zero(t, s.Len(), "Stack should stay at zero if popped without values")

	s.Push(1)
	assert.False(t, s.IsEmpty(), "Stack should not be empty if added value")
	assert.Equal(t, 1, s.Peek(), "Stack should return correct value when not empty")

	s.Push("string")
	assert.Equal(t, 2, s.Len(), "Stack should be 2 long")
	assert.Equal(t, "string", s.Peek(), "Stack should return correct value on top of stack")

	assert.Equal(t, "string", s.Pop(), "Stack should return the value that it pops")
	assert.Equal(t, 1, s.Pop(), "Stack should return the value that it pops")
	assert.Nil(t, s.Pop(), "Stack pop value should be nil when empty")
	assert.Nil(t, s.Peek(), "Stack peek value should be nil when empty")
	assert.Equal(t, 0, s.Len(), "Stack should be empty")
}

func TestIterStack(t *testing.T) {
	s1 := stack.New()

	s1.Push("3")
	s1.Push("2")
	s1.Push("1")

	for i, n := range s1.Items() {
		assert.Equal(t, fmt.Sprintf("%d", i+1), n)
	}

	s2 := stack.New()

	for _, n := range s2.Items() {
		assert.Nil(t, n)
		assert.Zero(t, n)
		fmt.Print("Should not print")
	}
}
