package array

import "testing"

func TestAnswer(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	expectResult := 10
	if expectResult != Answer1(heights) {
		t.Error()
	}
}
