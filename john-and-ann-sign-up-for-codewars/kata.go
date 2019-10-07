package kata

func johnAnn(n int) ([]int, []int) {
	j := []int{0, 0}
	a := []int{1, 1}

	for i := 2; i < n; i++ {
		tJ := j[i-1]
		j = append(j, i-a[tJ])

		tA := a[i-1]
		a = append(a, i-j[tA])
	}

	return j, a
}

func Ann(n int) []int {
	_, list := johnAnn(n)
	return list
}

func John(n int) []int {
	list, _ := johnAnn(n)
	return list
}

func SumJohn(n int) int {
	johns := John(n)

	sum := 0
	for _, val := range johns {
		sum += val
	}
	return sum
}

func SumAnn(n int) int {
	anns := Ann(n)

	sum := 0
	for _, val := range anns {
		sum += val
	}
	return sum
}
