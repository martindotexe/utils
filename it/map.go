package it

import "iter"

func Map[T, T2 any](i iter.Seq[T], f func(T) T2) iter.Seq[T2] {
	return func(yield func(T2) bool) {
		for value := range i {
			if !yield(f(value)) {
				return
			}
		}
	}
}

func MapError[T, T2 any](i iter.Seq[T], f func(T) (T2, error)) iter.Seq2[T2, error] {
	return func(yield func(T2, error) bool) {
		for value := range i {
			if !yield(f(value)) {
				return
			}
		}
	}
}
