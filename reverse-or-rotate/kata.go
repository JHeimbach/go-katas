package kata

import (
	"math"
	"strconv"
	"strings"
)

func Revrot(s string, n int) string {
	if n == 0 {
		return ""
	}

	var parts []string
	var part = ""
	for i, r := range []rune(s) {
		part += string(r)
		if (i+1)%n == 0 {
			parts = append(parts, part)
			part = ""
		}
	}

	modifiedParts := make([]string, len(parts))
	for i, num := range parts {
		sum := sumCubesOfDigits(num)

		if sum%2 == 0 {
			modifiedParts[i] = reverse(num)
		} else {
			modifiedParts[i] = moveN(num, 1)
		}
	}

	var returnString = ""
	for _, part := range modifiedParts {
		returnString += part
	}

	return returnString
}

func sumCubesOfDigits(num string) int {
	sum := 0.0
	for _, digit := range strings.Split(num, "") {
		digitInt, _ := strconv.Atoi(digit)
		sum += math.Pow(float64(digitInt), 3)
	}
	return int(sum)
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

func moveN(s string, n int) string {
	i := len(s)
	runes := make([]rune, i)
	a := []rune(s)
	for fi := 0; fi < i; fi++ {
		runes[fi] = a[(fi+n)%i]
	}
	return string(runes[:])
}
