package problem_6

import (
	"os"
	"testing"
)

func TestEditDistance(t *testing.T) {
	result, err := GetEditDistance([]byte("this is a test"), []byte("wokka wokka!!!"))

	if err != nil {
		t.Fatal("an error occured while getting distance", err)
	}
	if result != 37 {
		t.Fatal("wrong result", result)
	}
}

func TestGetKeyCandidatesSucceeded(t *testing.T) {
	content, err := os.ReadFile("6.txt")
	if err != nil {
		t.Fatal("error on opening file", err)
	}

	candidates, err := GetKeyCandidates(content, 2, 37, 3)
	if err != nil {
		t.Fatal("error on getting candidates", err)
	}

	t.Log("candidates:", candidates)
}
