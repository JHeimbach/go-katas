package kata

import (
	"strings"
)

func SpinWords(str string) string {

	words := strings.Fields(str)
	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverse(word)
		}
	}

	return strings.Join(words, " ")
}

func reverse(str string) string {
	revStr := make([]uint8, len(str))
	for i := 0; i < (len(str)+1)/2; i++ {
		revStr[i], revStr[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	return string(revStr)
}
