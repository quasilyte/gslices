package gslices

func Swap[T any](s []T, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Clone[T any](s []T) []T {
	return append(s[:0:0], s...)
}

func Transfer[T comparable](from, to *[]T, x T) {
	i := Index(*from, x)
	if i == -1 {
		panic("transfering a non-existing object")
	}
	DeleteAt(from, i)
	*to = append(*to, x)
}

func Last[T any](s []T) T {
	return s[len(s)-1]
}

func Prepend[T any](s *[]T, x T) {
	*s = append([]T{x}, *s...)
}
