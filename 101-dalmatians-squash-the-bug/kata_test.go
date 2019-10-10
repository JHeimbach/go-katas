package kata

import "testing"

func TestHowManyDalmatians(t *testing.T) {
	tests := []struct {
		name, want string
		input      int
	}{
		{name: "26 dogs", input: 26, want: DalmatiansMoreThanHandful},
		{name: "8 dogs", input: 8, want: DalmatiansHardlyAny},
		{name: "14 dogs", input: 14, want: DalmatiansMoreThanHandful},
		{name: "80 dogs", input: 80, want: DalmatiansLotsOfDogs},
		{name: "100 dogs", input: 100, want: DalmatiansLotsOfDogs},
		{name: "50 dogs", input: 50, want: DalmatiansMoreThanHandful},
		{name: "10 dogs", input: 10, want: DalmatiansHardlyAny},
		{name: "101 dogs", input: 101, want: Dalmatians101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HowManyDalmatians(tt.input)
			if got != tt.want {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}
