package gslices

func Push[T any](s *[]T, x T) {
	*s = append(*s, x)
}

func Pop[T any](s *[]T) T {
	l := len(*s)
	elem := (*s)[l-1]
	*s = (*s)[:l-1]
	return elem
}
