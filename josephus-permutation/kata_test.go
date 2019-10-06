package kata

import (
	"reflect"
	"testing"
)

func TestJosephus(t *testing.T) {
	type args struct {
		items []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "up to 10 with 1",
			args: args{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				k:     1,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "up to 10 with 2",
			args: args{
				items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				k:     2,
			},
			want: []int{2, 4, 6, 8, 10, 3, 7, 1, 9, 5},
		},
		{
			name: "up to 7 with 3",
			args: args{
				items: []int{1, 2, 3, 4, 5, 6, 7},
				k:     3,
			},
			want: []int{3, 6, 2, 7, 5, 1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Josephus(tt.args.items, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Josephus() = %v, want %v", got, tt.want)
			}
		})
	}
}
