package problem_4

import (
	"bufio"
	"os"
	"testing"
)

func TestProblem4(t *testing.T) {
	file, err := os.Open("4.txt")
	if err != nil {
		t.Fatal("failed to open file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := ""
	maxScore := 0

	for scanner.Scan() {
		decrypted, score, err := Decrypt(scanner.Text(), nil)
		if err != nil {
			t.Fatal("an error occured while decrypting", err)
		}
		if score > maxScore {
			maxScore = score
			result = decrypted
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatal("error occured while scanning file", err)
	}

	t.Log("successfully decrypted", result)
}
