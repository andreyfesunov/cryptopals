package problem_1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func HexToBase64(hs string) (string, error) {
	bytes, err := hex.DecodeString(hs)

	if err != nil {
		return "", fmt.Errorf("error occured while decoding hex string: %v", err)
	}

	log.Printf("%s", bytes)

	return base64.StdEncoding.EncodeToString(bytes), nil
}
