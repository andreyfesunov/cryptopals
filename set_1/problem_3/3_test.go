package problem_3

import "testing"

func TestProblem3(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	decrypted, err := Decrypt(input, nil)
	if err != nil {
		t.Fatal("an error occured while decrypting", err)
	}
	t.Log("successfully decrypted", decrypted)
}
