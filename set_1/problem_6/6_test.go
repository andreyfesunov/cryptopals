package problem_6

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"sort"
	"strings"
	"testing"

	xorkey "github.com/andreyfesunov/cryptopals/set_1/problem_3"
	encryption "github.com/andreyfesunov/cryptopals/set_1/problem_5"
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

	decoded, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		t.Fatal("error on decoding from base64", err)
	}

	candidates, err := GetKeyCandidates(decoded, 2, 37, 3)
	if err != nil {
		t.Fatal("error on getting candidates", err)
	}

	t.Log("candidates:", candidates)
}

func TestDecode(t *testing.T) {
	content, err := os.ReadFile("6.txt")
	if err != nil {
		t.Fatal("error on opening file", err)
	}

	decoded, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		t.Fatal("error on decoding from base64", err)
	}

	candidates, err := GetKeyCandidates(decoded, 2, 37, 3)
	if err != nil {
		t.Fatal("error on getting candidates", err)
	}

	results := make([]string, 0, len(candidates))

	for _, keysize := range candidates {
		blocks := GetBlocks(decoded, keysize)
		transposed := TransposeBlocks(blocks, keysize)
		var key strings.Builder
		for _, block := range transposed {
			decryptionResult, err := xorkey.Decrypt(hex.EncodeToString(block), nil)
			if err != nil {
				t.Fatal("error on decrypting block", err)
			}
			key.WriteRune(decryptionResult.Key)
		}
		answer, err := encryption.Encrypt(string(decoded), key.String())
		if err != nil {
			t.Fatal("error on encryption", err)
		}
		decodedResult, err := hex.DecodeString(answer)
		if err != nil {
			t.Fatal("error on decoding from hex", err)
		}
		results = append(results, string(decodedResult))
	}

	sort.Slice(results, func(i, j int) bool {
		return xorkey.ScoreEnglish(results[i], nil) > xorkey.ScoreEnglish(results[j], nil)
	})

	t.Log("result:", results[0])
}
