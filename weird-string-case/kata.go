package kata

import (
	"strings"
	"unicode"
)

func toWeirdCase(str string) string {
	inputWords := strings.Fields(strings.ToLower(str))

	var outputWords = make([]string, 0, len(inputWords))
	for _, word := range inputWords {
		outputWords = append(outputWords, string(weirdCaseWord(word)))
	}

	return strings.Join(outputWords, " ")
}

func weirdCaseWord(word string) []rune {
	oWord := make([]rune, 0, len(word))
	for i, letter := range []rune(word) {
		if i%2 == 0 {
			letter = unicode.ToUpper(letter)
		}
		oWord = append(oWord, letter)
	}
	return oWord
}
