package problem_6

import (
	"fmt"
	"math/bits"
	"sort"
)

func GetEditDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("slices should have equal lengths")
	}

	score := 0

	for i := range a {
		xor_result := a[i] ^ b[i]
		score += bits.OnesCount(uint(xor_result))
	}

	return score, nil
}

type ComparationResult struct {
	length int
	value  float64
}

func GetKeyCandidates(bytes []byte, minLength, maxLength, candidates int) ([]int, error) {
	results := make([]ComparationResult, 0, maxLength-minLength+1)

	for length := minLength; length <= maxLength; length += 1 {
		distance, err := GetEditDistance(bytes[:length], bytes[length:2*length])

		if err != nil {
			return nil, err
		}

		results = append(results, ComparationResult{
			length: length,
			value:  float64(distance) / float64(length),
		})
	}

	sort.Slice(results, func(a, b int) bool {
		return results[a].value < results[b].value
	})

	candidates = min(candidates, len(results))
	answers := make([]int, 0, candidates)

	for _, item := range results[:candidates] {
		answers = append(answers, item.length)
	}

	return answers, nil
}
