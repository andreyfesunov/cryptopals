package problem_2

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestProblem2(t *testing.T) {
	expected, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")

	a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	result, err := FixedXOR(a, b)

	if err != nil {
		t.Fatal("an error occured while applying fixed xor method", err)
	}
	if !bytes.Equal(result, expected) {
		t.Fatal("wrong result", result)
	}
}
