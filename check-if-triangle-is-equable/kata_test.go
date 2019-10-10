package kata

import (
	"fmt"
	"testing"
)

func TestEquableTriangle(t *testing.T) {
	tests := []struct {
		inputs [3]int
		want   bool
	}{
		{inputs: [3]int{5, 12, 13}, want: true},
		{inputs: [3]int{2, 3, 4}, want: false},
		{inputs: [3]int{6, 8, 10}, want: true},
		{inputs: [3]int{7, 15, 20}, want: true},
		{inputs: [3]int{17, 17, 30}, want: false},
		{inputs: [3]int{7, 10, 12}, want: false},
		{inputs: [3]int{6, 11, 12}, want: false},
		{inputs: [3]int{25, 25, 45}, want: false},
		{inputs: [3]int{13, 37, 30}, want: false},
		{inputs: [3]int{6, 25, 29}, want: true},
		{inputs: [3]int{10, 11, 18}, want: false},
		{inputs: [3]int{73, 9, 80}, want: false},
		{inputs: [3]int{12, 35, 37}, want: false},
		{inputs: [3]int{120, 109, 13}, want: false},
		{inputs: [3]int{9, 10, 17}, want: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should return %v", tt.inputs, tt.want), func(t *testing.T) {
			got := EquableTriangle(tt.inputs[0], tt.inputs[1], tt.inputs[2])

			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
