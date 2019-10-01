package kata

import (
	"strings"
)

func High(s string) string {
	highScore := 0
	highestWord := ""

	for _, word := range strings.Fields(s) {
		score := calcWordScore(word)
		if score > highScore {
			highScore = score
			highestWord = word
		}
	}

	return highestWord
}

func calcWordScore(word string) int {
	sum := 0
	for _, r := range strings.ToLower(word) {
		sum += int(r-'a') + 1
	}

	return sum
}
