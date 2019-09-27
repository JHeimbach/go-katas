package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 {
		return ""
	}
	inventoryMap := make(map[string]int)

	for _, letter := range listCat {
		inventoryMap[letter] = 0
	}

	for _, article := range listArt {
		parts := strings.Split(article, " ")
		firstLetter := []rune(parts[0])[0]
		num, _ := strconv.Atoi(parts[1])

		inventoryMap[string(firstLetter)] += num
	}
	var outputStrings []string

	for _, letter := range listCat {
		outputStrings = append(outputStrings, fmt.Sprintf("(%s : %d)", letter, inventoryMap[letter]))
	}

	return strings.Join(outputStrings, " - ")
}
