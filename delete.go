package gslices

func Delete[T comparable](s []T, x T) []T {
	if i := Index(s, x); i >= 0 {
		return DeleteAt(s, i)
	}
	return s
}

func DeleteAt[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
