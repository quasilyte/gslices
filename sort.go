package gslices

import (
	"slices"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](s []T) {
	slices.Sort(s)
}

func SortFunc[T any](slice []T, less func(a, b T) bool) {
	slices.SortFunc(slice, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})
}

func SortStableFunc[T any](s []T, less func(a, b T) bool) {
	slices.SortStableFunc(s, func(a, b T) int {
		if less(a, b) {
			return -1
		}
		if less(b, a) {
			return 1
		}
		return 0
	})
}
