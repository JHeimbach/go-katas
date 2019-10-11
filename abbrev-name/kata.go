package kata

import "strings"

func AbbrevName(name string) string {
	names := strings.Fields(strings.ToTitle(name))

	letters := []string{
		string(names[0][0]),
		string(names[len(names)-1][0]),
	}

	return strings.Join(letters, ".")
}
