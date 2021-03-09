package unclassified

import "testing"

func TestProblem001(t *testing.T) {
	var nums []int
	var expectResult, result int
	nums = []int{0, 1, 2, 3}
	expectResult = 4
	result = Problem001(nums)
	if result != expectResult {
		t.Errorf("nums=%v, expect_result=%d, result=%d", nums, expectResult, result)
	} else {
		t.Logf("nums=%v, result=%d", nums, result)
	}

	nums = []int{1, 2, 3}
	expectResult = 0
	result = Problem001(nums)
	if result != expectResult {
		t.Errorf("nums=%v, expect_result=%d, result=%d", nums, expectResult, result)
	} else {
		t.Logf("nums=%v, result=%d", nums, result)
	}

	nums = []int{0, 1, 2, 1}
	expectResult = 3
	result = Problem001(nums)
	if result != expectResult {
		t.Errorf("nums=%v, expect_result=%d, result=%d", nums, expectResult, result)
	} else {
		t.Logf("nums=%v, result=%d", nums, result)
	}

	nums = []int{1, 1, 1, 1}
	expectResult = 0
	result = Problem001(nums)
	if result != expectResult {
		t.Errorf("nums=%v, expect_result=%d, result=%d", nums, expectResult, result)
	} else {
		t.Logf("nums=%v, result=%d", nums, result)
	}
}
