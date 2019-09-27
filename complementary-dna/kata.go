package kata

import "strings"

var dnaReplacer = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
)

func DNAStrand(dna string) string {
	return dnaReplacer.Replace(dna)
}
