package kata

import (
	"fmt"
	"testing"
)

func TestOrderWeight(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "103 123 4444 99 2000",
			want:  "2000 103 123 4444 99",
		},
		{
			input: "2000 10003 1234000 44444444 9999 11 11 22 123",
			want:  "11 11 2000 10003 22 123 1234000 44444444 9999",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("OrderWeight(%v)", tt.input), func(t *testing.T) {
			if got := OrderWeight(tt.input); got != tt.want {
				t.Errorf("OrderWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
