package tests

import (
	"slices"
	"testing"

	"github.com/martindotexe/utils/it"
	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	add := func(i, j int) int { return i + j }

	sum := it.Fold(slices.Values([]int{1, 2, 3, 4}), add, 0)

	assert.Equal(t, 10, sum)

	mult := func(i, j int) int { return i * j }

	sum = it.Fold(slices.Values([]int{10, 5, 2}), mult, 1)

	assert.Equal(t, 100, sum)
}
