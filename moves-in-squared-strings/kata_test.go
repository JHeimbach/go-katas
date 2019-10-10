package kata

import "testing"

func TestVertMirror(t *testing.T) {
	t.Run("should handle basic cases Oper VertMirror", func(t *testing.T) {
		input := "hSgdHQ\nHnDMao\nClNNxX\niRvxxH\nbqTVvA\nwvSyRu"
		got := Oper(VertMirror, input)
		want := "QHdgSh\noaMDnH\nXxNNlC\nHxxvRi\nAvVTqb\nuRySvw"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestHorMirror(t *testing.T) {

	t.Run("should handle basic cases Oper HorMirror", func(t *testing.T) {
		input := "lVHt\nJVhv\nCSbg\nyeCt"
		got := Oper(HorMirror, input)
		want := "yeCt\nCSbg\nJVhv\nlVHt"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
