package kata

import "testing"

func TestIsValidIp(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "12.255.56.1",
			want: true,
		},
		{
			in:   "",
			want: false,
		},
		{
			in:   "abc.def.ghi.jkl",
			want: false,
		},
		{
			in:   "123.456.789.0",
			want: false,
		},
		{
			in:   "12.34.56",
			want: false,
		},
		{
			in:   "12.34.56 .1",
			want: false,
		},
		{
			in:   "12.34.56.-1",
			want: false,
		},
		{
			in:   "127.1.1.0",
			want: true,
		},
		{
			in:   "0.0.0.0",
			want: true,
		},
		{
			in:   "0.34.82.53",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := IsValidIp(tt.in); got != tt.want {
				t.Errorf("IsValidIp() = %v, want %v", got, tt.want)
			}
		})
	}
}
