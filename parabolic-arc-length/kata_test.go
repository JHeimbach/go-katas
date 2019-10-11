package kata

import (
	"fmt"
	"math"
	"testing"
)

func TestLenCurve(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  float64
	}{
		{
			input: 1,
			want:  1.4142133562,
		},
		{
			input: 2,
			want:  1.4604048132,
		},
		{
			input: 10,
			want:  1.478197397,
		},
		{
			input: 40,
			want:  1.478896272,
		},
		{
			input: 200,
			want:  1.478940994,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("parabolic length with %d steps", tt.input), func(t *testing.T) {
			assertFuzzyEquals(t, LenCurve(tt.input), tt.want)
		})
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("hypot parabolic length with %d steps ", tt.input), func(t *testing.T) {
			assertFuzzyEquals(t, LenCurveWithHypot(tt.input), tt.want)
		})
	}
}

func assertFuzzyEquals(t *testing.T, got float64, want float64) {
	t.Helper()

	var inrange bool
	merr := 1e-6
	var e float64
	e = math.Abs((got - want) / want)
	inrange = e <= merr
	if inrange == false {
		t.Errorf("Expected should be near: %1.6e , but got: %1.6e\n", want, got)
	}
}
