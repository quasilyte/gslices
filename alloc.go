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

// ResizeZeroed forces the slice s to be of the specified length.
// If s has enough capacity, the memory will be re-used.
// It does zeroing on the result.
func ResizeZeroed[T any](s *[]T, length int) {
	tmp := *s
	if cap(tmp) >= length {
		tmp = tmp[:length]
		clear(tmp)
	} else {
		tmp = make([]T, length)
	}
	*s = tmp
}

// Resize is like ResizeZeroed, but does no zeroing.
// Use it when you are sure that up to size elements
// will be overwritten right away (or you don't care about those).
func Resize[T any](s *[]T, length int) {
	tmp := *s
	if cap(tmp) >= length {
		tmp = tmp[:length]
	} else {
		tmp = make([]T, length)
	}
	*s = tmp
}
