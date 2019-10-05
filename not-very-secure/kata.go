package kata

import "unicode"

func alphanumeric(str string) bool {
	if str == "" {
		return false
	}
	for _, letter := range str {
		if !unicode.IsLetter(letter) && !unicode.IsDigit(letter) {
			return false
		}
	}
	return true
}
