package it

import (
	"iter"
)

func Reverse[V any](i func(func(V) bool)) iter.Seq[V] {
	l := Collect(i)
	for i, j := 0, len(l)-1; i < j; {
		l[i], l[j] = l[j], l[i]
		i++
		j--
	}

	return func(yield func(V) bool) {
		for _, v := range l {
			if !yield(v) {
				return
			}
		}
	}
}
