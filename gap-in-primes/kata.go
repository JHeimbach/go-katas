package kata

import (
	"math/big"
)

func Gap(g, m, n int) []int {
	prev := 0
	for curr := m; curr < n; curr++ {
		if !isPrime(curr) {
			continue
		}
		if curr-prev == g {
			return []int{prev, curr}
		}
		prev = curr
	}
	return nil
}

func isPrime(i int) bool {
	return big.NewInt(int64(i)).ProbablyPrime(0)
}
