package it

import "iter"

func Fold[T, T2 any](i iter.Seq[T], f func(T2, T) T2, acc T2) T2 {
	for v := range i {
		acc = f(acc, v)
	}
	return acc
}
