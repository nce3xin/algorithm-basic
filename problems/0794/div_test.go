package main

import (
	"fmt"
	"strconv"
	"testing"
)

type testCase struct {
	a    string
	b    int
	want *want
}

type want struct {
	c string // 商
	r int    // 余数
}

func TestDiv(t *testing.T) {
	testCases := []testCase{
		{
			a:    "7319951122",
			b:    19,
			want: &want{"385260585", 7},
		},
	}
	for _, tc := range testCases {
		A := Str2intArr(tc.a)
		b := tc.b
		C, r := div(A, b)
		got := ""
		for i := len(C) - 1; i >= 0; i-- {
			got += fmt.Sprintf("%d", C[i])
		}
		if got != tc.want.c || r != tc.want.r {
			t.Errorf("div(%v, %v), expected (%v, %v), got (%v,%v)",
				tc.a, tc.b, tc.want.c, tc.want.r, got, r)
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
