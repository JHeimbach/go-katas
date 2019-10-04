package kata

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  string
	}{
		{
			name:  "-6 to 20",
			input: []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20},
			want:  "-6,-3-1,3-5,7-11,14,15,17-20",
		},
		{
			name:  "-20 to 22",
			input: []int{-20, -15, -14, -13, -11, -10, -9, -1, 3, 10, 11, 12, 17, 21, 22},
			want:  "-20,-15--13,-11--9,-1,3,10-12,17,21,22",
		},
		{
			name:  "-34 to 11",
			input: []int{-34, -33, -29, -20, -14, -13, -12, -11, -10, -9, -8, -7, -5, -4, 0, 8, 10, 11},
			want:  "-34,-33,-29,-20,-14--7,-5,-4,0,8,10,11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.input); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
