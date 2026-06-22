package problem_2

import "fmt"

func FixedXOR(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("arrays lengths are not equal")
	}

	result := make([]byte, len(a))

	for i := range a {
		result[i] = a[i] ^ b[i]
	}

	return result, nil
}
