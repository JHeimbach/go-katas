package kata

import (
	"math"
)

func Solution(str string) []string {
	nPairs := int(math.Ceil(float64(len(str)) / 2))
	result := make([]string, 0, nPairs)

	// if string has odd length, add _ to make it even
	if len(str)%2 != 0 {
		str += "_"
	}

	for i := 0; i < nPairs; i++ {
		begin := i * 2
		result = append(result, str[begin:begin+2])
	}

	return result
}
