package kata

import (
	"fmt"
	"testing"
)

func TestSpinWords(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "Welcome",
			want:  "emocleW",
		},
		{
			input: "to",
			want:  "to",
		},
		{
			input: "CodeWars",
			want:  "sraWedoC",
		},
		{
			input: "Hey fellow warriors",
			want:  "Hey wollef sroirraw",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("SpinWords(%s) returns %v", tt.input, tt.want), func(t *testing.T) {
			if got := SpinWords(tt.input); got != tt.want {
				t.Errorf("SpinWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
