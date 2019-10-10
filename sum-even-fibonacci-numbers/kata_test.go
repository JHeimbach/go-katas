package kata

import "testing"

func TestSumEvenFibonacci(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{
			name:  "should return 10 for input 8",
			input: 8,
			want:  10,
		},
		{
			name:  "should return 60696 for input 111111",
			input: 111111,
			want:  60696,
		},
		{
			name:  "should return 4613732 for input 8675309",
			input: 8675309,
			want:  4613732,
		},
		{
			name:  "should return 2 for input 1",
			input: 1,
			want:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumEvenFibonacci(tt.input)

			if got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}
