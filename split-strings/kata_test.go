package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "abc",
			want:  []string{"ab", "c_"},
		},
		{
			input: "abcdef",
			want:  []string{"ab", "cd", "ef"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Solution(%s) returns %v", tt.input, tt.want), func(t *testing.T) {
			if got := Solution(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
