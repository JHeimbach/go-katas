package kata

func JosephusSurvivor(n, k int) int {

	items := makeItems(n)

	selector := k - 1
	for len(items) > 1 {
		if selector >= len(items) {
			selector = selector % len(items)
		}
		items = cutItem(items, selector)
		selector += k - 1
	}

	return items[0]
}

func makeItems(n int) []int {
	items := make([]int, n)
	for i := 0; i < len(items); i++ {
		items[i] = i + 1
	}
	return items
}

func cutItem(items []int, pos int) []int {
	return append(items[:pos], items[pos+1:]...)
}

func JosephusSurvivorRecursive(n, k int) int {
	if n == 1 {
		return 1
	}

	return (JosephusSurvivorRecursive(n-1, k)+k-1)%n + 1
}
