package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProductFib(t *testing.T) {
	tests := []struct {
		input uint64
		want  [3]uint64
	}{
		{
			input: 4895,
			want:  [3]uint64{55, 89, 1},
		},
		{
			input: 5895,
			want:  [3]uint64{89, 144, 0},
		},
		{
			input: 74049690,
			want:  [3]uint64{6765, 10946, 1},
		},
		{
			input: 84049690,
			want:  [3]uint64{10946, 17711, 0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ProductFib(%d) returns %v", tt.input, tt.want), func(t *testing.T) {
			if got := ProductFib(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductFib() = %v, want %v", got, tt.want)
			}
		})
	}
}
