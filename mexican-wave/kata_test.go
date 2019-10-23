package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWave(t *testing.T) {

	var tests = []struct {
		input string
		want  []string
	}{
		{
			input: " x yz",
			want:  []string{" X yz", " x Yz", " x yZ"},
		},
		{
			input: "abc",
			want:  []string{"Abc", "aBc", "abC"},
		},
		{
			input: "abc",
			want:  []string{"Abc", "aBc", "abC"},
		},
		{
			input: " ab  c",
			want:  []string{" Ab  c", " aB  c", " ab  C"},
		},
		{
			input: "",
			want:  []string{},
		},
		{
			input: "z",
			want:  []string{"Z"},
		},
		{
			input: "a a a a a",
			want:  []string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"},
		},
		{
			input: "aaaaa",
			want:  []string{"Aaaaa", "aAaaa", "aaAaa", "aaaAa", "aaaaA"},
		},
		{
			input: "                                                           ",
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("wave(%s) returns %v", tt.input, tt.want), func(t *testing.T) {
			got := wave(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
