package kata

func PositiveSum(numbers []int) (sum int) {
	for _, value := range numbers {
		if value > 0 {
			sum += value
		}
	}
	return sum
}
