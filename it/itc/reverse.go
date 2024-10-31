package itc

import "github.com/martindotexe/utils/it"

func (i Iterator[V]) Reverse() Iterator[V] {
	return Iterator[V](it.Reverse(i))
}
