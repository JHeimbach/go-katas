package kata

import "math/big"

func BackwardsPrime(start int, stop int) []int {
	var backwardsPrimes []int
	for i := start; i <= stop; i++ {
		if i != reverseInt(i) && isPrime(i) && isPrime(reverseInt(i)) {
			backwardsPrimes = append(backwardsPrimes, i)
		}
	}
	return backwardsPrimes
}

func isPrime(i int) bool {
	return big.NewInt(int64(i)).ProbablyPrime(0)
}

func reverseInt(input int) int {
	reversed := 0
	for input > 0 {
		remainder := input % 10
		reversed *= 10
		reversed += remainder
		input /= 10
	}
	return reversed
}
