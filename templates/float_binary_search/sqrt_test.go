package floatbinarysearch

import (
	"math"
	"testing"
)

type testCase struct {
	input float64
	want  float64
}

func TestSqrt(t *testing.T) {
	testCases := []testCase{
		{
			input: 3.0,
			want:  1.73205,
		},
		{
			input: 4.0,
			want:  2.00000,
		},
		{
			input: 0,
			want:  0.00000,
		},
		{
			input: 2,
			want:  1.41421,
		},
	}

	for _, tc := range testCases {
		got := sqrt(tc.input)
		if math.Abs(got-tc.want) > 1e-4 {
			t.Errorf("sqrt(%g), expected %g, got %g",
				tc.input, tc.want, got)
		}
	}
}
