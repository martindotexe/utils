package tests

import (
	"slices"
	"testing"

	"github.com/martindotexe/utils/it"
	"github.com/martindotexe/utils/it/filter"
	"github.com/martindotexe/utils/it/itc"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	compare := it.Collect(it.Filter(slices.Values([]int{1, 2, 3, 4, 5, 6, 7}), filter.IsEven))
	assert.Equal(t, []int{2, 4, 6}, compare)

	compare = it.Collect(it.Exclude(slices.Values([]int{1, 2, 3, 4, 5, 6, 7}), filter.IsEven))
	assert.Equal(t, []int{1, 3, 5, 7}, compare)

	// Test chained filter
	compare = itc.From(slices.Values([]int{1, 2, 3, 4, 5, 6, 7})).Filter(filter.IsEven).Collect()
	assert.Equal(t, []int{2, 4, 6}, compare)

	compare = itc.From(slices.Values([]int{1, 2, 3, 4, 5, 6, 7})).Exclude(filter.IsEven).Collect()
	assert.Equal(t, []int{1, 3, 5, 7}, compare)
}
