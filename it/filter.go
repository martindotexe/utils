package it

import "iter"

func Filter[T any](i iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range i {
			if f(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

func Exclude[T any](i iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return Filter(i, not(f))
}

func not[T any](f func(T) bool) func(T) bool {
	return func(value T) bool {
		return !f(value)
	}
}
