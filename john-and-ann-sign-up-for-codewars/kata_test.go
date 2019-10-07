package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestJohn(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{
			input: 11,
			want:  []int{0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6},
		},
		{
			input: 14,
			want:  []int{0, 0, 1, 2, 2, 3, 4, 4, 5, 6, 6, 7, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("John(%d) returns %v", tt.input, tt.want), func(t *testing.T) {

			got := John(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestAnn(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{
			input: 6,
			want:  []int{1, 1, 2, 2, 3, 3},
		},
		{
			input: 15,
			want:  []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 8, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Ann(%d) returns %v", tt.input, tt.want), func(t *testing.T) {

			got := Ann(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestSumAnn(t *testing.T) {
	want := 4070
	got := SumAnn(115)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumJohn(t *testing.T) {
	want := 1720
	got := SumJohn(75)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
