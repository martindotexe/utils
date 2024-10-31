package itc

import (
	"github.com/martindotexe/utils/it"
)

func (i Iterator[V]) Filter(f func(V) bool) Iterator[V] {
	return Iterator[V](it.Filter(i, f))
}

func (i Iterator[V]) Exclude(f func(V) bool) Iterator[V] {
	return Iterator[V](it.Exclude(i, f))
}
