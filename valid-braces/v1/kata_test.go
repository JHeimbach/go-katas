package v1

import (
	"fmt"
	"testing"
)

func TestValidBracesComplex(t *testing.T) {

	input := "(){[{}]}{}[]{}{}"
	want := true

	t.Run(fmt.Sprintf("%q is %q", input, "valid"), func(t *testing.T) {
		if got := ValidBraces(input); got != want {
			t.Errorf("ValidBraces() = %v, want %v", got, want)
		}
	})
}

func TestValidBraces(t *testing.T) {
	nameMap := map[bool]string{
		true:  "valid",
		false: "invalid",
	}

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "(){}[]",
			want:  true,
		},
		{
			input: "(}",
			want:  false,
		},
		{
			input: "()",
			want:  true,
		},
		{
			input: "[(])",
			want:  false,
		},
		{
			input: "[({)](]",
			want:  false,
		},
		{
			input: "([{}])",
			want:  true,
		},
		{
			input: "([]{[]})",
			want:  true,
		},
		{
			input: "{}({})[]",
			want:  true,
		},
		{
			input: "[]{()}([]()[])",
			want:  true,
		},
		{
			input: "(({[]{()}}[{(())}[]]))",
			want:  true,
		},
		{
			input: "(())[][]()[[]]",
			want:  true,
		},
		{
			input: "(){[{}]}{}[]{}{}",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q is %q", tt.input, nameMap[tt.want]), func(t *testing.T) {
			if got := ValidBraces(tt.input); got != tt.want {
				t.Errorf("ValidBraces() = %v, want %v", got, tt.want)
			}
		})
	}
}
