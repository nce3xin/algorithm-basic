package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	input []int
	want  []int
}

func TestQuickSort(t *testing.T) {
	testCases := []testCase{
		{
			input: []int{3, 1, 2, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			input: []int{},
			want:  []int{},
		},
		{
			input: []int{1},
			want:  []int{1},
		},
		{
			input: []int{1, 1},
			want:  []int{1, 1},
		},
		{
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		q := tc.input
		n := len(q)
		quickSort(q, 0, n-1)
		if !reflect.DeepEqual(q, tc.want) {
			t.Errorf("quickSort(%v): expected %v, got %v",
				tc.input, tc.want, q)
		}
	}
}
