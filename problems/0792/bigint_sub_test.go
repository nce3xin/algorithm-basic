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

func Str2intArr(s string) []int {
	var C []int
	for i := len(s) - 1; i >= 0; i-- {
		c, _ := strconv.Atoi(string(s[i]))
		C = append(C, c)
	}
	return C
}

func TestSub(t *testing.T) {
	testCases := []testCase{
		{
			a:    "113",
			b:    "110",
			want: "3",
		},
		{
			a:    "0",
			b:    "0",
			want: "0",
		},
		{
			a:    "12",
			b:    "35",
			want: "-23",
		},
		{
			a:    "35",
			b:    "12",
			want: "23",
		},
	}

	for _, tc := range testCases {
		A := Str2intArr(tc.a)
		B := Str2intArr(tc.b)
		got := ""
		if cmp(A, B) { // 如果A >= B
			C := sub(A, B)
			for i := len(C) - 1; i >= 0; i-- {
				got += fmt.Sprintf("%d", C[i])
			}
		} else {
			C := sub(B, A)
			got += fmt.Sprintf("-")
			for i := len(C) - 1; i >= 0; i-- {
				got += fmt.Sprintf("%d", C[i])
			}
		}
		if got != tc.want {
			t.Errorf("sub(%v,%v), expected %v, got %v",
				tc.a, tc.b, tc.want, got)
		}
	}
}
