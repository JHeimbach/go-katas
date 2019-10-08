package kata

import "strings"

func FirstNonRepeating(input string) string {
	if len(input) <= 1 {
		return input
	}

	characters := make(map[string]int)
	for _, r := range input {
		characters[strings.ToLower(string(r))]++
	}

	// loop over input again, because map has no guaranteed order
	for _, r := range input {
		count, _ := characters[strings.ToLower(string(r))] // ok can be omitted because every character is in map
		if count == 1 {
			return string(r)
		}
	}

	return ""
}
