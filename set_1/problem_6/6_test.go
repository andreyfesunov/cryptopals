package problem_6

import "testing"

func TestEditDistance(t *testing.T) {
	result, err := GetEditDistance([]byte("this is a test"), []byte("wokka wokka!!!"))

	if err != nil {
		t.Fatal("an error occured while getting distance", err)
	}
	if result != 37 {
		t.Fatal("wrong result", result)
	}
}
