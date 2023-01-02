package main

import (
	"fmt"
	"strconv"
	"testing"
)

type testCase struct {
	a    string
	b    string
	want string
}

func intArr2str(a []int) string {
	s := ""
	for i := len(a) - 1; i >= 0; i-- {
		s += fmt.Sprintf("%d", a[i])
	}
	return s
}

func str2intArr(s string) []int {
	var C []int
	for i := len(s) - 1; i >= 0; i-- {
		c, _ := strconv.Atoi(string(s[i]))
		C = append(C, c)
	}
	return C
}

func TestAdd(t *testing.T) {
	testCases := []testCase{
		{
			a:    "12",
			b:    "23",
			want: "35",
		},
		{
			a:    "1111111111111111111111111",
			b:    "1111111111111111111111111",
			want: "2222222222222222222222222",
		},
	}

	for _, tc := range testCases {
		A := str2intArr(tc.a)
		B := str2intArr(tc.b)
		C := add(A, B)
		got := intArr2str(C)
		if got != tc.want {
			t.Errorf("add(%v,%v), expected %v, got %v",
				A, B, tc.want, got)
		}
	}
}
