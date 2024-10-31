package itc

import "github.com/martindotexe/utils/it"

func (i Iterator[V]) Transform(f func(V) V) Iterator[V] {
	return Iterator[V](it.Map(i, f))
}

func (i Iterator[V]) TransformError(f func(V) (V, error)) Iterator2[V, error] {
	return Iterator2[V, error](it.MapError(i, f))
}
