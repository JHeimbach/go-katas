package kata

import "strings"

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func VertMirror(s string) string {
	words := strings.Split(s, "\n")

	for i, word := range words {
		words[i] = strings.Join(reverse(strings.Split(word, "")), "")
	}

	return strings.Join(words, "\n")
}

func HorMirror(s string) string {
	words := strings.Split(s, "\n")

	words = reverse(words)

	return strings.Join(words, "\n")
}

func reverse(w []string) []string {
	for i := 0; i < len(w)/2; i++ {
		backIndex := len(w) - 1 - i
		w[i], w[backIndex] = w[backIndex], w[i]
	}

	return w
}
