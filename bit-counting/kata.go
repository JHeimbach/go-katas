package kata

import (
	"math/bits"
	"strconv"
	"strings"
)

func CountBits(in uint) int {
	return strings.Count(strconv.FormatInt(int64(in), 2), "1")
}

// even simpler

func CountBitsSimpler(in uint) int {
	return bits.OnesCount(in)
}
