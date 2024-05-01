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
