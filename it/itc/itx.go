package itc

import (
	"iter"
	"slices"

	"github.com/martindotexe/utils/it"
)

type (
	Iterator[V any] func(func(V) bool)

	Iterator2[K, V any] func(func(K, V) bool)
)

func From[V any](i func(func(V) bool)) Iterator[V] {
	return Iterator[V](i)
}

func (i Iterator[V]) Seq() iter.Seq[V] {
	return iter.Seq[V](i)
}

func (i Iterator2[K, V]) Seq() iter.Seq2[K, V] {
	return iter.Seq2[K, V](i)
}

func (iterator Iterator[V]) Collect() []V {
	return slices.Collect(iter.Seq[V](iterator))
}

func (iterator Iterator2[V, V2]) Collect() ([]V, []V2) {
	return it.Collect2(iterator)
}
