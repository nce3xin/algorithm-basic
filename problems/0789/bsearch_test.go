package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	arr  []int
	x    int
	want []int
}

func TestBinarySearch(t *testing.T) {
	testCases := []testCase{
		{
			arr:  []int{1, 2, 2, 3, 3, 4},
			x:    3,
			want: []int{3, 4},
		},
		{
			arr:  []int{1, 2, 2, 3, 3, 4},
			x:    4,
			want: []int{5, 5},
		},
		{
			arr:  []int{1, 2, 2, 3, 3, 4},
			x:    5,
			want: []int{-1, -1},
		},
	}

	for _, tc := range testCases {
		lb := binarySearchLowerBound(tc.arr, tc.x, 0, len(tc.arr)-1)
		ub := binarySearchUpperBound(tc.arr, tc.x, 0, len(tc.arr)-1)
		res := []int{lb, ub}
		if !reflect.DeepEqual(res, tc.want) {
			t.Errorf("BinarySearch(%v,%d), expected %v, got %v", tc.arr,
				tc.x, tc.want, res)
		}
	}
}
