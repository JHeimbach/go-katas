package kata

import (
	"testing"
)

func TestStockList(t *testing.T) {
	tests := []struct {
		name  string
		inArt []string
		inCat []string
		want  string
	}{
		{
			name:  "basic case",
			inArt: []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"},
			inCat: []string{"A", "B", "C", "D"},
			want:  "(A : 0) - (B : 1290) - (C : 515) - (D : 600)",
		},
		{
			name:  "some more",
			inArt: []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"},
			inCat: []string{"A", "B"},
			want:  "(A : 200) - (B : 1140)",
		},
		{
			name:  "empty art",
			inArt: []string{},
			inCat: []string{"A", "B"},
			want:  "",
		},
		{
			name:  "nil art",
			inCat: []string{"A", "B"},
			want:  "",
		},
		{
			name:  "empty cat",
			inArt: []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"},
			inCat: []string{},
			want:  "",
		},
		{
			name:  "nil cat",
			inArt: []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"},
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StockList(tt.inArt, tt.inCat)
			if got != tt.want {
				t.Errorf("Stocklist(%v, %v) returned %q, expected %q", tt.inArt, tt.inCat, got, tt.want)
			}
		})
	}
}
