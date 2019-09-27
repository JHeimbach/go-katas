package kata

import "testing"

func TestCreatePhoneNumber(t *testing.T) {
	expected := "(123) 456-7890"
	result := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})

	if result != expected {
		t.Errorf("got %q, want %q", result, expected)
	}
}
