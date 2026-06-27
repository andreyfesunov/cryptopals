package problem_8

import (
	"bufio"
	"encoding/hex"
	"os"
	"testing"
)

func TestProblem8(t *testing.T) {
	file, err := os.Open("8.txt")
	if err != nil {
		t.Fatal("error on opening file", err)
	}
	defer file.Close()

	reader := bufio.NewScanner(file)

	maxScore, result := 0, ""

	for reader.Scan() {
		content := reader.Text()
		bytes, err := hex.DecodeString(content)
		if err != nil {
			t.Fatal("error on decoding content", err)
		}
		if len(bytes)%16 != 0 {
			t.Fatal("error on bytes invariant - can't split into 16-length blocks")
		}

		seen := make(map[[16]byte]bool)
		duplicates := 0

		for start := 0; start < len(bytes); start += 16 {
			var block [16]byte
			copy(block[:], bytes[start:start+16])
			if seen[block] {
				duplicates++
			} else {
				seen[block] = true
			}
		}

		if maxScore < duplicates {
			maxScore = duplicates
			result = content
		}
	}

	t.Log("result:", result)
}
