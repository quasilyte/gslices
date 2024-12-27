package gslices

func Move[T any](s []T, from, to int) []T {
	if from == to {
		return s
	}
	// TODO: this can probably be done with a single append
	// operation instead of doing it twice in DeleteAt
	// and then in InsertAt.
	// But before this function shows up in the
	// profiler, I'm not going to bother.
	x := s[from]
	DeleteAt(&s, from)
	s = InsertAt(s, x, to)
	return s
}
