package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	parts := strings.Fields(in)
	lowest, _ := strconv.Atoi(parts[0])
	highest := lowest

	for _, part := range parts {
		number, _ := strconv.Atoi(part)
		if number < lowest {
			lowest = number
		}
		if number > highest {
			highest = number
		}
	}
	return fmt.Sprintf("%d %d", highest, lowest)
}
