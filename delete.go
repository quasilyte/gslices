package gslices

func Delete[T comparable](s []T, x T) []T {
	if i := Index(s, x); i >= 0 {
		return DeleteAt(s, i)
	}
	return s
}

func DeleteFunc[T any](s []T, f func(x T) bool) []T {
	if i := IndexFunc(s, f); i >= 0 {
		return DeleteAt(s, i)
	}
	return s
}

func DeleteAt[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}
