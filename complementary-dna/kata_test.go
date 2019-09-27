package kata

import "testing"

func TestDNAStrand(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "AAAA",
			want: "TTTT",
		},
		{
			in:   "ATTGC",
			want: "TAACG",
		},
		{
			in:   "GTAT",
			want: "CATA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := DNAStrand(tt.in); got != tt.want {
				t.Errorf("DNAStrand() = %q, want %q", got, tt.want)
			}
		})
	}
}
