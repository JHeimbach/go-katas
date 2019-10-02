package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDirReduc(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"},
			want:  []string{"WEST"},
		},
		{
			input: []string{"NORTH", "WEST", "SOUTH", "EAST"},
			want:  []string{"NORTH", "WEST", "SOUTH", "EAST"},
		},
		{
			input: []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"},
			want:  []string{"NORTH"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("DirReduc(%v)", tt.input), func(t *testing.T) {
			if got := DirReduc(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DirReduc() = %q, want %q", got, tt.want)
			}
		})
	}
}
