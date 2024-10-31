package it

import "iter"

func Enumerate[T any](i iter.Seq[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		count := 0

		for value := range i {
			if !yield(count, value) {
				return
			}
			count++
		}
	}
}
