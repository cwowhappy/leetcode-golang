package dp

/**
Problem300 最长上升子序列
时间复杂度: O(n^2)
空间复杂度: O(n)
 */
func LengthOfLIS(nums []int) int {
	maxLength := 0
	numsLength := len(nums)
	if numsLength > 0 {
		maxLength = 1
		lengths := make([]int, 0, numsLength)
		for i := 0; i < numsLength; i ++ {
			lengths = append(lengths, 1)
		}
		for i := 0; i < numsLength; i ++ {
			for j := i; j < numsLength; j ++ {
				if nums[j] > nums[i] && lengths[i] + 1 > lengths[j] {
					lengths[j] = lengths[i] + 1
					if lengths[j] > maxLength {
						maxLength = lengths[j]
					}
				}
			}
		}
	}
	return maxLength
}

/**
Problem300 最长上升子序列
时间复杂度: O(n * log(n))
空间复杂度: O(n)
*/
func LengthOfLISV2(nums []int) int {
	maxLength := 0
	numsLength := len(nums)
	if numsLength > 0 {
		maxLength = 1
		lengths := make([]int, 0, numsLength)
		for i := 0; i < numsLength; i ++ {
			lengths = append(lengths, 1)
		}
		for i := 0; i < numsLength; i ++ {
			for j := i; j < numsLength; j ++ {
				if nums[j] > nums[i] && lengths[i] + 1 > lengths[j] {
					lengths[j] = lengths[i] + 1
					if lengths[j] > maxLength {
						maxLength = lengths[j]
					}
				}
			}
		}
	}
	return maxLength
}
