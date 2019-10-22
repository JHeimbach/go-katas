package kata

import (
	"fmt"
	"testing"
)

func Test_toWeirdCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "abc def",
			want:  "AbC DeF",
		},
		{
			input: "ABC",
			want:  "AbC",
		},
		{
			input: "This is a test Looks like you passed",
			want:  "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("toWeirdCase(%q) returns %q", tt.input, tt.want), func(t *testing.T) {
			if got := toWeirdCase(tt.input); got != tt.want {
				t.Errorf("toWeirdCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
