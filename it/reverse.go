package it

import (
	"iter"
)

func Reverse[T any](i iter.Seq[T]) iter.Seq[T] {
	l := Collect(i)
	for i, j := 0, len(l)-1; i < j; {
		l[i], l[j] = l[j], l[i]
		i++
		j--
	}

	return func(yield func(T) bool) {
		for _, v := range l {
			if !yield(v) {
				return
			}
		}
	}
}
