package kata

import (
	"fmt"
	"testing"
)

func TestStrong(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{
			input: 1,
			want:  IsStrong,
		},
		{
			input: 2,
			want:  IsStrong,
		},
		{
			input: 145,
			want:  IsStrong,
		},
		{
			input: 7,
			want:  NotStrong,
		},
		{
			input: 3,
			want:  NotStrong,
		},
		{
			input: 93,
			want:  NotStrong,
		},
		{
			input: 185,
			want:  NotStrong,
		},
		{
			input: 40585,
			want:  IsStrong,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d should be %q", tt.input, tt.want), func(t *testing.T) {
			if got := Strong(tt.input); got != tt.want {
				t.Errorf("Strong() = %v, want %v", got, tt.want)
			}
		})
	}
}
