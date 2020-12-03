package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	n := 2
	result := 2
	if ClimbStairs(n) != result {
		t.Error()
	}
	n = 3
	result = 3
	if ClimbStairs(n) != result {
		t.Error()
	}
}
