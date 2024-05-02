package gslices

func Swap[T any](s []T, i, j int) {
	s[i], s[j] = s[j], s[i]
}
