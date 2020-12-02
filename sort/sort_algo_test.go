package sort_test

import (
	"github.com/cwowhappy/leetcode-golang/sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 2, 4, 1, 3}
	sortedNums := []int{1, 2, 3, 4, 5}
	sort.BubbleSort(nums)
	if !checkSortResult(sortedNums, nums) {
		t.Error(nums, sortedNums)
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 2, 4, 1, 3}
	sortedNums := []int{1, 2, 3, 4, 5}
	sort.MergeSort(nums)
	if !checkSortResult(sortedNums, nums) {
		t.Error(nums, sortedNums)
	}
}

func TestQuickSort(t *testing.T) {
	nums := []int{5, 2, 4, 1, 3}
	sortedNums := []int{1, 2, 3, 4, 5}
	sort.QuickSort(nums)
	if !checkSortResult(sortedNums, nums) {
		t.Error(nums, sortedNums)
	}
	nums = []int{5, 4, 3, 2, 1}
	sortedNums = []int{1, 2, 3, 4, 5}
	sort.QuickSort(nums)
	if !checkSortResult(sortedNums, nums) {
		t.Error(nums, sortedNums)
	}
}

func checkSortResult(sortedNums []int, resultNums []int) (flag bool) {
	flag = true
	sortedNumsLength := len(sortedNums)
	resultNumsLength := len(resultNums)

	if sortedNumsLength == resultNumsLength {
		for i := 0; i < sortedNumsLength; i++ {
			if sortedNums[i] != resultNums[i] {
				flag = false
				break
			}
		}
	} else {
		flag = false
	}
	return
}
