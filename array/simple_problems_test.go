package array

import "testing"

func TestMajorityElementHard(t *testing.T) {
	var nums []int
	var mainNum int
	nums = []int{1, 2}
	mainNum = -1
	if result := MajorityElementHard(nums); result != mainNum {
		t.Error(nums, mainNum, result)
	}

	nums = []int{1, 2, 1}
	mainNum = 1
	if result := MajorityElementHard(nums); result != mainNum {
		t.Error(nums, mainNum, result)
	}
}