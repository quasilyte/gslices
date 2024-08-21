package gslices

func Contains[T comparable](s []T, x T) bool {
	return Index(s, x) >= 0
}

func Index[T comparable](s []T, x T) int {
	for i := range s {
		if x == s[i] {
			return i
		}
	}
	return -1
}

func IndexFunc[T comparable](s []T, f func(x T) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}

func Find[T comparable](s []T, x T) (result T) {
	for i := range s {
		if x == s[i] {
			return s[i]
		}
	}
	return result
}

func FindFunc[T comparable](s []T, f func(x T) bool) (result T) {
	for i := range s {
		if f(s[i]) {
			return s[i]
		}
	}
	return result
}
