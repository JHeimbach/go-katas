package kata

import (
	"strconv"
	"strings"
)

func CountBits(in uint) int {
	return strings.Count(strconv.FormatInt(int64(in), 2), "1")
}
