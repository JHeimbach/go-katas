package kata

import (
	"math/big"
)

func EasyLine(n int) string {
	if n == 0 {
		return "1"
	}
	n64 := int64(n)

	sum := big.NewInt(n64).Binomial(n64*2, n64)

	return sum.String()
}
