package kata

func Josephus(items []int, k int) []int {
	result := make([]int, 0)

	selector := k - 1
	for len(items) > 0 {
		if selector >= len(items) {
			selector = selector % len(items)
		}
		result = append(result, items[selector])
		items = append(items[:selector], items[selector+1:]...)
		selector += k - 1
	}

	return result
}
