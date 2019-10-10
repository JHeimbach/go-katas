package v1

import (
	"strings"
)

var brackets = []string{
	"(",
	"{",
	"[",
}
var pairMap = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func ValidBraces(str string) bool {
	if str == "" {
		return true
	}
	if len(str) < 2 {
		return false
	}

	valid := false

	for _, bracket := range brackets {
		last := strings.LastIndex(str, bracket)
		if last < 0 {
			continue
		}
		next := strings.Index(str[last:], pairMap[bracket]) + last
		if next < last {
			continue
		}

		if next-last == 1 {
			valid = ValidBraces(str[:last] + str[next+1:])
		}

		if valid = ValidBraces(str[last+1 : next]); !valid {
			break
		}

		if valid = ValidBraces(str[:last] + str[next+1:]); !valid {
			break
		}
	}

	return valid
}
