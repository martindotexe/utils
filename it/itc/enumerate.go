package itc

import "github.com/martindotexe/utils/it"

func (i Iterator[V]) Enumerate() Iterator2[int, V] {
	return Iterator2[int, V](it.Enumerate(i))
}
