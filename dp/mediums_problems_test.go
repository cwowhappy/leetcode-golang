package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := 4
	if result != LengthOfLIS(nums) {
		t.Error(nums, result)
	}
}

func TestFindCheapestPrice(t *testing.T) {
	result := 200
	n, src, dst, k := 3, 0, 2, 1
	flights := [][]int{
		{0, 1, 100}, {1, 2, 100}, {0, 2, 500},
	}
	if result != FindCheapestPrice(n, flights, src, dst, k) {
		t.Error("error")
	}
}