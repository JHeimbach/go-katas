package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	results := []string{}

	for _, value := range array1 {
		for _, test := range array2 {
			if strings.Contains(test, value) && !sliceContains(results, value) {
				results = append(results, value)
			}
		}
	}

	sort.Strings(results)

	return results
}

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
