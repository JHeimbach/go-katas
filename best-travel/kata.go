package kata

type combinations struct {
	list    []int
	results [][]int
}

func (c *combinations) calc(r, index, i int, data []int) {
	if index == r {
		tmp := make([]int, len(data))
		copy(tmp, data)
		c.results = append(c.results, tmp)
		return
	}

	if i == len(c.list) {
		return
	}

	data[index] = c.list[i]
	c.calc(r, index+1, i+1, data)
	c.calc(r, index, i+1, data)
}

func allCombinations(list []int, k int) [][]int {
	combinations := combinations{
		list: list,
	}
	combinations.calc(k, 0, 0, make([]int, k))

	return combinations.results
}

func ChooseBestSum(t, k int, ls []int) int {
	if k > len(ls) {
		return -1
	}

	biggestSum := -1
	for _, set := range allCombinations(ls, k) {
		sum := sumUp(set)
		if sum > biggestSum && sum <= t {
			biggestSum = sum
		}
	}

	return biggestSum
}

func sumUp(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
