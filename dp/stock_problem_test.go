package dp

import "testing"

func TestProblem121(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expectResult := 5
	if expectResult != Problem121(prices) {
		t.Error()
	}
}

func TestProblem122(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expectResult := 7
	if expectResult != Problem122(prices) {
		t.Error()
	}
}