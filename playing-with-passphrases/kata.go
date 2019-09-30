package kata

import (
	"strings"
)

func PlayPass(s string, n int) string {
	s = strings.ToUpper(s)
	returnString := ""
	for i, letter := range []rune(s) {
		if letter >= 65 && letter <= 90 {
			letter = letter + rune(n)

			if letter > 90 {
				letter = (letter % 90) + 64
			}

			if (i+1)%2 == 0 {
				letter += 32
			}
		}

		if letter >= 48 && letter <= 57 {
			letter = (9 - (letter - 48)) + 48
		}

		returnString += string(letter)
	}

	return reverse(returnString)
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, r := range s {
		n--
		runes[n] = r
	}
	return string(runes[:])
}
