package main

import (
	"fmt"
	"strconv"
	"testing"
)

type testCase struct {
	a    string
	b    int
	want string
}

func TestMul(t *testing.T) {
	testCases := []testCase{
		{
			a:    "123",
			b:    12,
			want: "1476",
		},
		{
			a:    "12",
			b:    0,
			want: "0",
		},
	}

	for _, tc := range testCases {
		A := Str2intArr(tc.a)
		b := tc.b
		C := mul(A, b)
		got := ""
		for i := len(C) - 1; i >= 0; i-- {
			got += fmt.Sprintf("%d", C[i])
		}
		if got != tc.want {
			t.Errorf("mul(%v,%d), expected %v, got %v",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func Str2intArr(s string) []int {
	var C []int
	for i := len(s) - 1; i >= 0; i-- {
		c, _ := strconv.Atoi(string(s[i]))
		C = append(C, c)
	}
	return C
}
