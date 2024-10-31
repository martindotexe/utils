package it

import "iter"

func Enumerate[V any](i func(func(V) bool)) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		count := 0

		for value := range i {
			if !yield(count, value) {
				return
			}
			count++
		}
	}
}
