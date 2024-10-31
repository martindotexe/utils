package filter

import "golang.org/x/exp/constraints"

func IsEven[T constraints.Integer](integer T) bool {
	return integer%2 == 0
}

func IsOdd[T constraints.Integer](integer T) bool {
	return integer%2 != 0
}

func IsEqual[T comparable](value T) func(T) bool {
	return func(candidate T) bool {
		return value == candidate
	}
}
