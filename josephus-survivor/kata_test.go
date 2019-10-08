package kata

import (
	"fmt"
	"testing"
)

func TestJosephusSurvivor(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 7,
				k: 3,
			},
			want: 4,
		},
		{
			args: args{
				n: 11,
				k: 19,
			},
			want: 10,
		},
		{
			args: args{
				n: 40,
				k: 3,
			},
			want: 28,
		},
		{
			args: args{
				n: 2,
				k: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d with %d resulted %d", tt.args.n, tt.args.k, tt.want), func(t *testing.T) {
			if got := JosephusSurvivor(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("JosephusSurvivor() = %v, want %v", got, tt.want)
			}
			if got := JosephusSurvivorRecursive(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("JosephusSurvivorRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkJosephusSurvivor(b *testing.B) {
	for i := 1; i < b.N; i++ {
		JosephusSurvivor(i, 3)
	}
}

func BenchmarkJosephusSurvivorRecursive(b *testing.B) {
	for i := 1; i < b.N; i++ {
		JosephusSurvivorRecursive(i, 3)
	}
}
