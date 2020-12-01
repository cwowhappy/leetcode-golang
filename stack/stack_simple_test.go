package stack

import "testing"

func TestIsValid(t *testing.T) {
	s := "()"
	if !IsValid(s) {
		t.Error(s)
	}
}
