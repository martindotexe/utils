package tests

import (
	"slices"
	"testing"

	"github.com/martindotexe/utils/filter"
	"github.com/martindotexe/utils/it"
	"github.com/stretchr/testify/assert"
)

func FilterTest(t *testing.T) {
	even := []int{2, 4, 6}
	compare := []int{}

	for v := range it.Filter(slices.Values([]int{1, 2, 3, 4, 5, 6, 7}), filter.IsEven) {
		compare = append(compare, v)
	}

	assert.Equal(t, even, compare)
}
