package kata

import "strings"

func DNAStrand(dna string) (result string) {
	dna = strings.ToUpper(dna)

	complements := map[string]string{"A": "T", "T": "A", "C": "G", "G": "C"}

	for _, char := range dna {
		result += complements[string(char)]
	}
	return
}
