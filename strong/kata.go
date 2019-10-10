package kata

const IsStrong string = "STRONG!!!!"
const NotStrong string = "Not Strong !!"

func Strong(n int) string {
	digits := extractDigits(n)

	sum := 0
	for _, d := range digits {
		sum += factorial(d)
	}

	if sum == n {
		return IsStrong
	}

	return NotStrong
}

func extractDigits(n int) (digits []int) {
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	return digits
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	if n < 3 {
		return n
	}

	fac := 1
	for i := 2; i <= n; i++ {
		fac *= i
	}

	return fac
}
