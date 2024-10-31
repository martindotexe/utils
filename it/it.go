package it

func Collect[V any](i func(func(V) bool)) []V {
	c := []V{}

	for v := range i {
		c = append(c, v)
	}
	return c
}

func Collect2[K, V any](i func(func(K, V) bool)) ([]K, []V) {
	l := []K{}
	r := []V{}

	for k, v := range i {
		l = append(l, k)
		r = append(r, v)
	}
	return l, r
}

func TryCollect[V any](iterator func(func(V, error) bool)) ([]V, error) {
	var values []V

	for v, err := range iterator {
		if err != nil {
			return values, err
		}
		values = append(values, v)
	}

	return values, nil
}
