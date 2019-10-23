package kata

import (
	"strings"
	"unicode"
)

func wave(words string) []string {
	output := make([]string, 0, len(words))

	for key, letter := range words {
		if unicode.IsSpace(letter) {
			continue
		}
		output = append(output, words[:key]+strings.ToUpper(string(letter))+words[key+1:])
	}

	return output
}
