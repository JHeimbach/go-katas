package kata

import (
	"fmt"
	"testing"
)

func Test_alphanumeric(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{
			input: ".*?",
			want:  false,
		},
		{
			input: "a",
			want:  true,
		},
		{
			input: "Mazinkaiser",
			want:  true,
		},
		{
			input: "hello world_",
			want:  false,
		},
		{
			input: "PassW0rd",
			want:  true,
		},
		{
			input: "     ",
			want:  false,
		},
		{
			input: "",
			want:  false,
		},
		{
			input: "\n\t\n",
			want:  false,
		},
		{
			input: "ciao\n$$_",
			want:  false,
		},
		{
			input: "__ * __",
			want:  false,
		},
		{
			input: "&)))(((",
			want:  false,
		},
		{
			input: "43534h56jmTHHF3k",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%q should be %v", tt.input, tt.want), func(t *testing.T) {
			if got := alphanumeric(tt.input); got != tt.want {
				t.Errorf("alphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
