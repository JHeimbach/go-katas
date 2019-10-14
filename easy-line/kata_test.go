package kata

import (
	"fmt"
	"testing"
)

func TestEasyLine(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{0, "1"},
		{1, "2"},
		{7, "3432"},
		{13, "10400600"},
		{100, "90548514656103281165404177077484163874504589675413336841320"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("should handle basic case %d", tt.input), func(t *testing.T) {
			if got := EasyLine(tt.input); got != tt.want {
				t.Errorf("EasyLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
