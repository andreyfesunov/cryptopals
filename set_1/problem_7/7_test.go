package problem_7

import (
	"crypto/aes"
	"encoding/base64"
	"os"
	"testing"
)

const key = "YELLOW SUBMARINE"

func TestProblem7(t *testing.T) {
	content, err := os.ReadFile("7.txt")
	if err != nil {
		t.Fatal("error on reading file content", err)
	}

	decoded, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		t.Fatal("error on decoding content", err)
	}

	keyBytes := []byte(key)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		t.Fatal("error on creating new cipher", err)
	}

	decrypted := make([]byte, len(decoded))
	for i := 0; i < len(decoded); i += aes.BlockSize {
		block.Decrypt(decrypted[i:i+aes.BlockSize], decoded[i:i+aes.BlockSize])
	}

	t.Log("decrypted:", string(decrypted))
}
