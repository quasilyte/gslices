package gslices

func DeleteAt[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
