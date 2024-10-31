package tests

import (
	"slices"
	"testing"

	"github.com/martindotexe/utils/it"
	"github.com/martindotexe/utils/it/itc"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	double := func(v int) int { return v * 2 }
	compare := it.Collect(it.Map(slices.Values([]int{1, 2, 3, 4, 5}), double))
	assert.Equal(t, []int{2, 4, 6, 8, 10}, compare)

	doubleErr := func(v int) (int, error) { return v * 2, nil }
	compare, errs := it.Collect2(it.MapError(slices.Values([]int{1, 2, 3, 4, 5}), doubleErr))
	assert.Equal(t, []int{2, 4, 6, 8, 10}, compare)
	assert.Equal(t, []error{nil, nil, nil, nil, nil}, errs)

	// Test from transform
	compare = itc.From(slices.Values([]int{1, 2, 3, 4, 5})).
		Transform(double).
		Collect()
	assert.Equal(t, []int{2, 4, 6, 8, 10}, compare)

	compare, errs = itc.From(slices.Values([]int{1, 2, 3, 4, 5})).
		TransformError(doubleErr).
		Collect()
	assert.Equal(t, []int{2, 4, 6, 8, 10}, compare)
	assert.Equal(t, []error{nil, nil, nil, nil, nil}, errs)
}
