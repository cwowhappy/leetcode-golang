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

func TestMinCostClimbingStairs(t *testing.T) {
	costs := []int{0, 0, 1, 1}
	result := 1
	if result != MinCostClimbingStairs(costs) {
		t.Error()
	}
}