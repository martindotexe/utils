package tests

import (
	"testing"

	"github.com/martindotexe/utils/filter"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert.True(t, filter.IsEven(uint8(2)))
	assert.True(t, filter.IsEven(int16(-2)))
	assert.False(t, filter.IsEven(int32(2147483647)))
	assert.False(t, filter.IsEven(uint64(18446744073709551615)))

	assert.True(t, filter.IsOdd(int8(3)))
	assert.True(t, filter.IsOdd(uint32(417)))
	assert.True(t, filter.IsOdd(uint64(18446744073709551615)))
	assert.False(t, filter.IsOdd(uint64(18446744073709551612)))
}
