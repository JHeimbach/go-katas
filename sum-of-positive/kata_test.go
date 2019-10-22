package kata

import (
	"fmt"
	"testing"
)

func TestPositiveSum(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			input: []int{1, -2, 3, 4, 5},
			want:  13,
		},
		{
			input: []int{},
			want:  0,
		},
		{
			input: []int{-1, -2, -3, -4, -5},
			want:  0,
		},
		{
			input: []int{-1, 2, 3, 4, -5},
			want:  9,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("PositiveSum(%d) returns %d", tt.input, tt.want), func(t *testing.T) {
			if got := PositiveSum(tt.input); got != tt.want {
				t.Errorf("PositiveSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
