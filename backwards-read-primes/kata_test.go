package kata

import (
	"reflect"
	"testing"
)

func TestBackwardsPrime(t *testing.T) {
	t.Run("all cases between 1 and 100", func(t *testing.T) {
		var want = []int{13, 17, 31, 37, 71, 73, 79, 97}
		got := BackwardsPrime(1, 100)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("BackwardsPrime(): got %v wanted %v", got, want)
		}
	})
	t.Run("all cases between 1 and 31", func(t *testing.T) {
		var want = []int{13, 17, 31}
		got := BackwardsPrime(1, 31)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("BackwardsPrime(): got %v wanted %v", got, want)
		}
	})
	t.Run("all cases between 501 and 599", func(t *testing.T) {
		var want []int
		got := BackwardsPrime(501, 599)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("BackwardsPrime(): got %v wanted %v", got, want)
		}
	})
}
