package problem_6

import (
	"fmt"
	"math/bits"
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
