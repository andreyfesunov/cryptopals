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

type ComparisonResult struct {
	length int
	value  float64
}

func GetKeyCandidates(bytes []byte, minLength, maxLength, candidates int) ([]int, error) {
	results := make([]ComparisonResult, 0, maxLength-minLength+1)

	for length := minLength; length <= maxLength; length += 1 {
		if length*2 > len(bytes) {
			break
		}

		total, pairs := 0.0, 0

		for i := 0; i+2*length <= len(bytes); i += length {
			distance, err := GetEditDistance(bytes[i:i+length], bytes[i+length:i+2*length])

			if err != nil {
				return nil, err
			}

			total += float64(distance) / float64(length)
			pairs += 1
		}

		results = append(results, ComparisonResult{
			length: length,
			value:  total / float64(pairs),
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

func GetBlocks(ciphertext []byte, keysize int) [][]byte {
	var blocks [][]byte
	for start := 0; start < len(ciphertext); start += keysize {
		end := min(start+keysize, len(ciphertext))
		blocks = append(blocks, ciphertext[start:end])
	}
	return blocks
}

func TransposeBlocks(blocks [][]byte, keysize int) [][]byte {
	transposed := make([][]byte, keysize)
	for i := range keysize {
		for _, block := range blocks {
			if i < len(block) {
				transposed[i] = append(transposed[i], block[i])
			}
		}
	}
	return transposed
}
