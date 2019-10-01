package kata

import (
	"fmt"
	"testing"
)

func TestChooseBestSum(t *testing.T) {
	tests := []struct {
		list        []int
		maxDistance int
		maxCities   int
		want        int
	}{
		{
			list:        []int{50, 55, 56, 57, 58},
			maxDistance: 163,
			maxCities:   3,
			want:        163,
		},
		{
			list:        []int{50},
			maxDistance: 163,
			maxCities:   3,
			want:        -1,
		},
		{
			list:        []int{91, 74, 73, 85, 73, 81, 87},
			maxDistance: 230,
			maxCities:   3,
			want:        228,
		},
		{
			list:        []int{91, 74, 73, 85, 73, 81, 87},
			maxDistance: 331,
			maxCities:   2,
			want:        178,
		},
		{
			list:        []int{91, 74, 73, 85, 73, 81, 87},
			maxDistance: 331,
			maxCities:   4,
			want:        331,
		},
		{
			list:        []int{91, 74, 73, 85, 73, 81, 87},
			maxDistance: 331,
			maxCities:   5,
			want:        -1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("ChooseBestSum(%v) returns %d", tt.list, tt.want), func(t *testing.T) {
			got := ChooseBestSum(tt.maxDistance, tt.maxCities, tt.list)

			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
