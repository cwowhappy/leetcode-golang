package array

import "testing"

func TestLongestSubarray(t *testing.T) {
	nums := []int{8,2,4,7}
	limit := 4
	_ = LongestSubarray(nums, limit)
}
