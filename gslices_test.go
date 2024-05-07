package gslices

import (
	"reflect"
	"testing"
)

func TestMove(t *testing.T) {
	tests := []struct {
		s    []int
		from int
		to   int
		want []int
	}{
		{[]int{10, 20, 30}, 0, 0, []int{10, 20, 30}},
		{[]int{10, 20, 30}, 3, 3, []int{10, 20, 30}},

		{[]int{10, 20, 30, 40}, 3, 0, []int{40, 10, 20, 30}},
		{[]int{10, 20, 30, 40}, 3, 1, []int{10, 40, 20, 30}},
		{[]int{10, 20, 30, 40}, 3, 2, []int{10, 20, 40, 30}},

		{[]int{10, 20, 30, 40}, 0, 1, []int{20, 10, 30, 40}},
		{[]int{10, 20, 30, 40}, 0, 2, []int{20, 30, 10, 40}},
		{[]int{10, 20, 30, 40}, 0, 3, []int{20, 30, 40, 10}},

		{[]int{10, 20, 30, 40}, 1, 0, []int{20, 10, 30, 40}},
		{[]int{10, 20, 30, 40}, 1, 2, []int{10, 30, 20, 40}},
		{[]int{10, 20, 30, 40}, 1, 3, []int{10, 30, 40, 20}},
	}

	for _, test := range tests {
		have := Move(Clone(test.s), test.from, test.to)
		want := test.want
		if !reflect.DeepEqual(have, want) {
			t.Fatalf("move(%#v, %v, %v):\nhave: %#v\nwant: %#v", test.s, test.from, test.to, have, want)
		}
		have2 := Move(Clone(have), test.to, test.from)
		if !reflect.DeepEqual(have2, test.s) {
			t.Fatalf("second move(%#v, %v, %v):\nhave: %#v\nwant: %#v", have, test.to, test.from, have2, test.s)
		}
	}
}
