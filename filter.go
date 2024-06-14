package gslices

func FilterInplace[T any](s []T, f func(x T) bool) []T {
	keep := s[:0]
	for _, x := range s {
		if f(x) {
			keep = append(keep, x)
		}
	}
	return keep
}
