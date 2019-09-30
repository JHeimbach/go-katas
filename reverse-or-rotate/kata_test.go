package kata

import "testing"

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRevrot(t *testing.T) {
	t.Run("test empty n", func(t *testing.T) {
		want := ""
		got := Revrot("1234", 0)

		assertString(t, got, want)
	})

	t.Run("test str to short", func(t *testing.T) {
		want := ""
		got := Revrot("1234", 5)

		assertString(t, got, want)
	})

	t.Run("test long string", func(t *testing.T) {
		want := "234561876549"
		got := Revrot("123456987654", 6)

		assertString(t, got, want)
	})

	t.Run("test long string", func(t *testing.T) {
		want := "44668753"
		got := Revrot("66443875", 4)

		assertString(t, got, want)
	})

	t.Run("test long string", func(t *testing.T) {
		want := "330479108928157"
		got := Revrot("733049910872815764", 5)

		assertString(t, got, want)
	})
}
