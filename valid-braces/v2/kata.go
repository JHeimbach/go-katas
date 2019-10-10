package v2

// after looking at some solutions

import (
	"strings"
)

var pairMap = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func ValidBraces(str string) bool {
	stack := make([]string, 0)

	for _, c := range strings.Split(str, "") {
		switch c {
		case "(", "{", "[":
			stack = append(stack, pairMap[c])
		case ")", "}", "]":
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != c {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
