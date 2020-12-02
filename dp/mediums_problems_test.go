package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := 4
	if result != LengthOfLIS(nums) {
		t.Error(nums, result)
	}
}