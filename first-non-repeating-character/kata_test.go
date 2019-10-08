package kata

import (
	"fmt"
	"testing"
)

func TestFirstNonRepeating(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "a",
			want:  "a",
		},
		{
			input: "stress",
			want:  "t",
		},
		{
			input: "moonmen",
			want:  "e",
		},
		{
			input: "",
			want:  "",
		},
		{
			input: "abba",
			want:  "",
		},
		{
			input: "aa",
			want:  "",
		},
		{
			input: "~><#~><",
			want:  "#",
		},
		{
			input: "hello world, eh?",
			want:  "w",
		},
		{
			input: "sTreSS",
			want:  "T",
		},
		{
			input: "Go hang a salami, I'm a lasagna hog!",
			want:  ",",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("with %q returns %q", tt.input, tt.want), func(t *testing.T) {
			got := FirstNonRepeating(tt.input)
			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
