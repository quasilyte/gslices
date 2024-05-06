package gslices

func InsertAt[T any](s []T, x T, i int) []T {
	return append(s[:i], append([]T{x}, s[i:]...)...)
}
