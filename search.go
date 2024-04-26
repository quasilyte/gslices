package gslices

func Contains[T comparable](s []T, v T) bool {
	return Index(s, v) >= 0
}

func Index[T comparable](s []T, v T) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}
