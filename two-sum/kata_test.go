package kata

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		args args
		want [2]int
	}{
		{
			args: args{numbers: []int{1, 2, 3}, target: 4},
			want: [2]int{0, 2},
		},
		{
			args: args{numbers: []int{1234, 5678, 9012}, target: 14690},
			want: [2]int{1, 2},
		},
		{
			args: args{numbers: []int{2, 2, 4}, target: 4},
			want: [2]int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v with %d returns %v", tt.args.numbers, tt.args.target, tt.want), func(t *testing.T) {
			if got := TwoSum(tt.args.numbers, tt.args.target); got != tt.want {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
