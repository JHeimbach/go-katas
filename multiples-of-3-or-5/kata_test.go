package kata

import "testing"

func TestMultiple3And5(t *testing.T) {
	want := 23
	got := Multiple3And5(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
