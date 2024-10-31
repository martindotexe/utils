package tests

import (
	"slices"
	"testing"

	"github.com/martindotexe/utils/filter"
	"github.com/martindotexe/utils/it/itc"
	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	l, r := itc.From(slices.Values([]int{1, 2, 3, 4, 5})).
		Filter(filter.IsEven).
		Transform(func(i int) int { return i * 10 }).
		Reverse().
		Enumerate().
		Collect()

	assert.Equal(t, []int{0, 1}, l)
	assert.Equal(t, []int{40, 20}, r)
}
