package problem_5

import "testing"

func TestProblem5(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"

	output := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`

	result, err := Encrypt(input, key)

	if err != nil {
		t.Fatal("an error occured while encrypting string", err)
	}
	if output != result {
		t.Fatal("wrong encryption", result)
	}
}
