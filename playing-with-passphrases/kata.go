package kata

import (
	"strings"
	"unicode"
)

func PlayPass(s string, n int) string {
	s = strings.ToUpper(s)

	returnString := ""
	for i, letter := range s {

		if unicode.IsLetter(letter) {
			letter = 'A' + (letter-'A'+int32(n))%26
		}

		if unicode.IsDigit(letter) {
			letter = '9' - letter + '0'
		}

		if i%2 == 1 && unicode.IsLetter(letter) {
			letter = unicode.ToLower(letter)
		}

		returnString = string(letter) + returnString
	}

	return returnString
}
