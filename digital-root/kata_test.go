package kata

import (
	"fmt"
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 16,
			want:  7,
		},
		{
			input: 195,
			want:  6,
		},
		{
			input: 992,
			want:  2,
		},
		{
			input: 167346,
			want:  9,
		},
		{
			input: 0,
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("DigitalRoot(%d) returns %d", tt.input, tt.want), func(t *testing.T) {
			if got := DigitalRoot(tt.input); got != tt.want {
				t.Errorf("DigitalRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
