package it

import "iter"

func Collect[T any](i iter.Seq[T]) []T {
	c := []T{}

	for v := range i {
		c = append(c, v)
	}
	return c
}
