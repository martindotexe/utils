package it

func Fold[V, A any](i func(func(V) bool), f func(A, V) A, acc A) A {
	for v := range i {
		acc = f(acc, v)
	}
	return acc
}
