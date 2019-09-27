package kata

import (
	"testing"
)

func TestHighAndLow(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "the solution returns the correct value",
			in:   "8 3 -5 42 -1 0 0 -9 4 7 4 -4",
			want: "42 -9",
		},
		{
			name: "single input",
			in:   "42",
			want: "42 42",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighAndLow(tt.in); got != tt.want {
				t.Errorf("HighAndLow() = %v, want %v", got, tt.want)
			}
		})
	}
}
