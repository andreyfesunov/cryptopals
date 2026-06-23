package problem_4

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
)

func ScoreEnglish(s string, goodChars []string) int {
	score := 0
	if goodChars == nil {
		goodChars = []string{"a", "e", "i", "o", "u", "r", "s", "t", "l", "m", "n", " "}
	}
	for _, item := range goodChars {
		score += strings.Count(strings.ToLower(s), item)
	}
	return score
}

func GetXORFunction(i rune) func(rune) rune {
	return func(r rune) rune {
		return r ^ i
	}
}

func Decrypt(encrypted string, goodChars []string) (string, int, error) {
	encryptedBytes, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", 0, fmt.Errorf("an error occured while decoding string: %v", err)
	}

	maxScore := 0
	var result []byte

	for i := range 256 {
		text := bytes.Map(GetXORFunction(rune(i)), encryptedBytes)
		score := ScoreEnglish(string(text), goodChars)
		if score > maxScore {
			maxScore = score
			result = text
		}
	}

	return string(result), maxScore, nil
}
