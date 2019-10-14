package kata

import (
	"math"
)

func Simpson(n int) float64 {
	b := math.Pi
	h := b / float64(n)
	nHalfs := float64(n / 2)

	result := 2 * sum(nHalfs-1, func(i float64) float64 {
		return f(2 * i * h)
	})

	result += 4 * sum(nHalfs, func(i float64) float64 {
		return f(((2 * i) - 1) * h)
	})

	result += f(b)

	result *= b / float64(3*n)

	return result
}

func sum(to float64, f func(i float64) float64) float64 {
	sum := 0.0
	for i := 1.0; i <= to; i++ {
		sum += f(i)
	}
	return sum
}

func f(x float64) float64 {
	return 1.5 * math.Pow(math.Sin(x), 3)
}
