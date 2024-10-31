package tests

import (
	"slices"
	"testing"

	"github.com/martindotexe/utils/it"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	reverse := it.Reverse(slices.Values([]int{1, 2, 3}))

	c := slices.Collect(reverse)

	assert.Equal(t, []int{3, 2, 1}, c)
}
