package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGap(t *testing.T) {
	type args struct {
		g int
		m int
		n int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				g: 2,
				m: 100,
				n: 110,
			},
			want: []int{101, 103},
		},
		{
			args: args{
				g: 2,
				m: 5,
				n: 5,
			},
			want: nil,
		},
		{
			args: args{
				g: 4,
				m: 100,
				n: 110,
			},
			want: []int{103, 107},
		},
		{
			args: args{
				g: 4,
				m: 130,
				n: 200,
			},
			want: []int{163, 167},
		},
		{
			args: args{
				g: 6,
				m: 100,
				n: 110,
			},
			want: nil,
		},
		{
			args: args{
				g: 8,
				m: 300,
				n: 400,
			},
			want: []int{359, 367},
		},
		{
			args: args{
				g: 10,
				m: 300,
				n: 400,
			},
			want: []int{337, 347},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("step %d between %d, %d", tt.args.g, tt.args.m, tt.args.n), func(t *testing.T) {
			if got := Gap(tt.args.g, tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gap() = %v, want %v", got, tt.want)
			}
		})
	}
}
