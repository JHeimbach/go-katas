package kata

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  PosPeaks
	}{
		{
			name:  "should support finding peaks",
			input: []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			want:  PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			name:  "should support finding peaks, but should ignore peaks on the edge of the array",
			input: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			want:  PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		},
		{
			name:  "should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau",
			input: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			want:  PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		},
		{
			name:  "should support finding peaks, but should ignore peaks on the edge of the array",
			input: []int{2, 1, 3, 1, 2, 2, 2, 2},
			want:  PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		},
		{
			name:  "2 14 13 -5 -5 6 4 9 0 1 1",
			input: []int{2, 14, 13, -5, -5, 6, 4, 9, 0, 1, 1},
			want:  PosPeaks{Pos: []int{1, 5, 7}, Peaks: []int{14, 6, 9}},
		},
		{
			name:  "-3 12 12 9 -3 14 5 1 -4",
			input: []int{-3, 12, 12, 9, -3, 14, 5, 1, -4},
			want:  PosPeaks{Pos: []int{1, 5}, Peaks: []int{12, 14}},
		},
		{
			name:  "should return an object with empty arrays if the input does not contain any peak",
			input: []int{},
			want:  PosPeaks{Pos: []int{}, Peaks: []int{}},
		},
		{
			name:  "3 -5 14 11 11 -4 12 12 7",
			input: []int{3, -5, 14, 11, 11, -4, 12, 12, 7},
			want:  PosPeaks{Pos: []int{2, 6}, Peaks: []int{14, 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PickPeaks(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PickPeaks() = %v, want %v", got, tt.want)
			}
		})
	}
}
