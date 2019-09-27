package kata

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	t.Run("should handle basic cases", func(t *testing.T) {
		want := []string{"arp", "live", "strong"}
		a1 := []string{"live", "arp", "strong"}
		a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}

		got := InArray(a1, a2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("InArray() returns %v want %v", got, want)
		}
	})
	t.Run("should handle basic cases", func(t *testing.T) {
		want := []string{}
		a1 := []string{"cod", "code", "wars", "ewar", "ar"}
		a2 := []string{}
		got := InArray(a1, a2)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("InArray() returns %+v want %v", got, want)
		}
	})
}
