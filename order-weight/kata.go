package kata

import (
	"sort"
	"strings"
)

type WeightWithScore struct {
	weight string
	score  int
}

type WeightsWithScore []WeightWithScore

func (w WeightsWithScore) Len() int {
	return len(w)
}

func (w WeightsWithScore) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WeightsWithScore) Less(i, j int) bool {
	if w[i].score == w[j].score {
		return w[i].weight < w[j].weight
	}
	return w[i].score < w[j].score
}

func (w WeightsWithScore) String() string {
	var output []string
	for _, w := range w {
		output = append(output, w.weight)
	}

	return strings.Join(output, " ")
}

func OrderWeight(input string) string {
	listWeights := strings.Fields(input)
	weights := make(WeightsWithScore, len(listWeights))

	for i, weight := range listWeights {
		weights[i] = WeightWithScore{
			weight: weight,
			score:  calcScore(weight),
		}
	}

	sort.Sort(weights)

	return weights.String()
}

func calcScore(weight string) int {
	score := 0
	for _, digit := range weight {
		score += int(digit - '0')
	}
	return score
}
