package gslices

func AllocElem[T any](s *[]T) *T {
	if len(*s) >= cap(*s) {
		var empty T
		*s = append(*s, empty)
		return &(*s)[len(*s)-1]
	}

	*s = (*s)[:len(*s)+1]
	return &(*s)[len(*s)-1]
}

// ReallocZeroed forces the slice s to be of the specified size.
// If s has enough capacity, the memory will be re-used.
// It does zeroing on the result.
func ReallocZeroed[T any](s *[]T, size int) {
	tmp := *s
	switch {
	case cap(tmp) == 0:
		tmp = make([]T, size)
	case cap(tmp) >= size:
		tmp = tmp[:size]
		clear(tmp)
	default:
		// Re-use and zero max amount of memory, and then
		// append some more to meet the required size.
		tmp = tmp[:cap(tmp)]
		clear(tmp)
		tmp = append(tmp, make([]T, size-cap(tmp))...)
	}
	*s = tmp
}

// Realloc is like ReallocZeroed, but does no zeroing.
// Use it when you are sure that up to size elements
// will be overwritten right away (or you don't care about those).
func Realloc[T any](s *[]T, size int) {
	tmp := *s
	switch {
	case cap(tmp) == 0:
		tmp = make([]T, size)
	case cap(tmp) >= size:
		tmp = tmp[:size]
	default:
		tmp = tmp[:cap(tmp)]
		tmp = append(tmp, make([]T, size-cap(tmp))...)
	}
	*s = tmp
}
