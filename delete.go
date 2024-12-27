package gslices

func Delete[T comparable](s *[]T, x T) {
	if i := Index(*s, x); i >= 0 {
		DeleteAt(s, i)
	}
}

func DeleteFunc[T any](s *[]T, f func(x T) bool) {
	if i := IndexFunc(*s, f); i >= 0 {
		DeleteAt(s, i)
	}
}

func DeleteAt[T any](s *[]T, i int) {
	*s = append((*s)[:i], (*s)[i+1:]...)
}
