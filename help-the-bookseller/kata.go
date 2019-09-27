package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	inventoryMap := make(map[string]int)

	for _, article := range listArt {
		parts := strings.Split(article, " ")
		firstLetter := parts[0][:1]
		num, _ := strconv.Atoi(parts[1])
		inventoryMap[firstLetter] += num
	}

	var outputStrings []string
	for _, letter := range listCat {
		outputStrings = append(outputStrings, fmt.Sprintf("(%s : %d)", letter, inventoryMap[letter]))
	}

	return strings.Join(outputStrings, " - ")
}
