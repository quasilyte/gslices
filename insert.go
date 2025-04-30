package gslices

func InsertAt[T any](s []T, x T, i int) []T {
	return append(s[:i], append([]T{x}, s[i:]...)...)
}

func InsertSliceAt[T any](s []T, s2 []T, i int) []T {
	result := s[:i]
	result = append(result, s2...)
	result = append(result, s[i:]...)
	return result
}
