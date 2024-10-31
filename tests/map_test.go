package tests

import (
	"errors"
	"slices"
	"strconv"
	"testing"

	"github.com/martindotexe/utils/it"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func MapTest(t *testing.T) {
	var nums string
	for v := range it.Map(slices.Values([]int{1, 2, 3}), strconv.Itoa) {
		nums += v
	}
	assert.Equal(t, "123", nums)
}

func MapErrorTest(t *testing.T) {
	nums := []int{}

	for v, err := range it.MapError(slices.Values([]string{"1", "2", "3"}), strconv.Atoi) {
		nums = append(nums, v)
		require.NoError(t, err)
	}
	assert.Equal(t, []int{1, 2, 3}, nums)

	alwaysError := func(v int) (int, error) {
		return 0, errors.New("Always error")
	}

	for _, err := range it.MapError(slices.Values([]int{1, 2, 3}), alwaysError) {
		require.Error(t, err)
	}
}
