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

}
