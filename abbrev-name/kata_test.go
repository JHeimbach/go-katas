package kata

import "testing"

func TestAbbrevName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "Sam Harris",
			want:  "S.H",
		},
		{
			input: "Patrick Feenan",
			want:  "P.F",
		},
		{
			input: "Evan Cole",
			want:  "E.C",
		},
		{
			input: "P Favuzzi",
			want:  "P.F",
		},
		{
			input: "David Mendieta",
			want:  "D.M",
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := AbbrevName(tt.input); got != tt.want {
				t.Errorf("AbbrevName() = %q, want %q", got, tt.want)
			}
		})
	}
}
