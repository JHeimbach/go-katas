package kata

func SumEvenFibonacci(limit int) int {
	sum, f1, f2 := 2, 1, 2

	for f2 < limit {
		f1, f2 = f2, f1+f2

		if f2%2 == 0 {
			sum += f2
		}
	}

	return sum
}
