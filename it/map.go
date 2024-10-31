package it

import "iter"

func Map[V, V2 any](i func(func(V) bool), f func(V) V2) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		for value := range i {
			if !yield(f(value)) {
				return
			}
		}
	}
}

func MapError[V, V2 any](i func(func(V) bool), f func(V) (V2, error)) iter.Seq2[V2, error] {
	return func(yield func(V2, error) bool) {
		for value := range i {
			if !yield(f(value)) {
				return
			}
		}
	}
}
