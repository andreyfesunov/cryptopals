package problem_5

import "encoding/hex"

func Encrypt(s, key string) (string, error) {
	bytes, keyBytes := []byte(s), []byte(key)
	result := make([]byte, len(bytes))

	for i := range bytes {
		result[i] = bytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	return hex.EncodeToString(result), nil
}
