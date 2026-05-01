package gslices

func AllocElem[T any](s *[]T) *T {
	if len(*s) >= cap(*s) {
		var empty T
		*s = append(*s, empty)
		return &(*s)[len(*s)-1]
	}

	*s = (*s)[:len(*s)+1]
	return &(*s)[len(*s)-1]
}
