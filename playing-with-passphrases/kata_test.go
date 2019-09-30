package kata

import "testing"

func TestPlayPass(t *testing.T) {
	t.Run("i love you with 1", func(t *testing.T) {
		want := "!!!vPz fWpM J"
		got := PlayPass("I LOVE YOU!!!", 1)

		assertString(t, got, want)
	})
	t.Run("i love you with 0", func(t *testing.T) {
		want := "!!!uOy eVoL I"
		got := PlayPass("I LOVE YOU!!!", 0)

		assertString(t, got, want)
	})
	t.Run("BORN IN 2015!", func(t *testing.T) {
		want := "!4897 Oj oSpC"
		got := PlayPass("BORN IN 2015!", 1)

		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
