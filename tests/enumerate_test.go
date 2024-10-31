package tests

import (
	"slices"
	"testing"

	"github.com/martindotexe/utils/it"
	"github.com/stretchr/testify/assert"
)

func TestEnumerate(t *testing.T) {
	k, v := it.Collect2(it.Enumerate(slices.Values([]int{1, 2, 3, 4, 5})))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, k)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, v)
}
