package kata

import (
	"fmt"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		input uint
		want  int
	}{
		{
			input: 0,
			want:  0,
		},
		{
			input: 4,
			want:  1,
		},
		{
			input: 7,
			want:  3,
		},
		{
			input: 9,
			want:  2,
		},
		{
			input: 10,
			want:  2,
		},
	}
	t.Run("CountBits with string conversion", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("CountBits(%d) returns %d", tt.input, tt.want), func(t *testing.T) {
				if got := CountBits(tt.input); got != tt.want {
					t.Errorf("CountBits() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("CountBits with math/bits", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(fmt.Sprintf("CountBitsSimpler(%d) returns %d", tt.input, tt.want), func(t *testing.T) {
				if got := CountBitsSimpler(tt.input); got != tt.want {
					t.Errorf("CountBitsSimpler() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
