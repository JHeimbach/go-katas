package kata

import (
	"fmt"
	"testing"
)

func TestHigh(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"man i need a taxi up to ubud", "taxi"},
		{"what time are we climbing up the volcano", "volcano"},
		{"take me to semynak", "semynak"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("High(%q) returns %q", tt.input, tt.want), func(t *testing.T) {
			got := High(tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
