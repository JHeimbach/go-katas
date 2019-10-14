package kata

import (
	"fmt"
	"math"
	"testing"
)

func TestSimpson(t *testing.T) {
	tests := []struct {
		input int
		want  float64
	}{
		{input: 290, want: 1.9999999986},
		{input: 72, want: 1.9999996367},
		{input: 252, want: 1.9999999975},
		{input: 40, want: 1.9999961668},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Simpson(%d) returns %f", tt.input, tt.want), func(t *testing.T) {
			got := Simpson(tt.input)
			fuzzyEquals(t, got, tt.want)
		})
	}
}

func fuzzyEquals(t *testing.T, got, want float64) {
	t.Helper()

	var mErr = 1e-10
	var e float64
	if want == 0.0 {
		e = math.Abs(got)
	} else {
		e = math.Abs((got - want) / want)
	}
	if e >= mErr {
		t.Errorf("Expected should be near: %1.10e , but got: %1.10e\n", want, got)
	}
}

func Test_f(t *testing.T) {
	tests := []struct {
		x    float64
		want float64
	}{
		{
			0, 0,
		},
		{
			math.Pi / 2, 3 / float64(2),
		},
		{
			math.Pi, 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("f(%f) returns %f", tt.x, tt.want), func(t *testing.T) {
			got := f(tt.x)
			fuzzyEquals(t, got, tt.want)
		})
	}
}
