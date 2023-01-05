package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	s    string
	want []string
}

func TestPrintWords(t *testing.T) {
	testCases := []testCase{
		{
			s:    "abc def ghi",
			want: []string{"abc", "def", "ghi"},
		},
	}

	for _, tc := range testCases {
		got := printWords(tc.s)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("printWords(%v), expected %v, got %v",
				tc.s, tc.want, got)
		}
	}
}
