package kata

func DigitalRoot(n int) int {
	if countDigits(n) < 2 {
		return n
	}

	sum := 0
	for _, digit := range toDigits(n) {
		sum += digit
	}

	return DigitalRoot(sum)
}

func countDigits(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func toDigits(n int) []int {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}
