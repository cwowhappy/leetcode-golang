package sort

func swapNums(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func BubbleSort(nums []int) {
	numsLength := len(nums)
	for i := 0; i < numsLength; i++ {
		for j := i + 1; j < numsLength; j++ {
			if nums[i] > nums[j] {
				swapNums(nums, i, j)
			}
		}
	}
	return
}

func MergeSort(nums []int) {
	recursionMergeSort(nums, 0, len(nums) - 1)
}

func recursionMergeSort(nums []int, indexBegin int, indexEnd int) {
	if indexBegin >= indexEnd {
		return
	}
	indexMid := indexBegin + (indexEnd - indexBegin) / 2
	recursionMergeSort(nums, indexBegin, indexMid)
	recursionMergeSort(nums, indexMid, indexEnd)

}

func QuickSort(nums []int) {
	recursionQuickSort(nums, 0, len(nums) - 1)
}

func recursionQuickSort(nums []int, indexBegin int, indexEnd int) {
	if indexBegin >= indexEnd {
		return
	}
	guardIndex := indexBegin
	guardValue := nums[guardIndex]
	i, j := indexBegin + 1, indexEnd
	for i <= j {
		for ; i <= j && guardValue >= nums[i]; i ++ {}
		for ; i <= j && guardValue < nums[j]; j -- {}
		if i < j {
			swapNums(nums, i, j)
		}
	}
	if guardIndex < j {
		swapNums(nums, guardIndex, j)
	}
	recursionQuickSort(nums, indexBegin, j - 1)
	recursionQuickSort(nums, j + 1, indexEnd)
}
