package kata

import (
	"fmt"
	"strings"
)

type numberRange struct {
	first, last int
}

func (n numberRange) String() string {
	if n.first == n.last {
		return fmt.Sprintf("%d", n.first)
	}
	return fmt.Sprintf("%d-%d", n.first, n.last)
}

func (n numberRange) isRange() bool {
	return n.first != n.last
}

func Solution(list []int) string {
	var ranges []numberRange
	var first, last = list[0], list[0]

	for _, value := range list[1:] {
		if last == value-1 {
			last = value
			continue
		}
		ranges = appendRange(numberRange{first, last}, ranges)

		first = value
		last = value
	}

	ranges = appendRange(numberRange{first, last}, ranges)

	output := make([]string, len(ranges))
	for i, r := range ranges {
		output[i] = r.String()
	}

	return strings.Join(output, ",")
}

func appendRange(n numberRange, ranges []numberRange) []numberRange {
	if n.last-n.first >= 2 {
		return append(ranges, n)
	}
	ranges = append(ranges, numberRange{n.first, n.first})

	if n.isRange() {
		ranges = append(ranges, numberRange{n.last, n.last})
	}

	return ranges
}
