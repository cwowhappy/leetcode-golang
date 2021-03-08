package array

func LongestSubarray(nums []int, limit int) int {
	if len(nums) == 0 || limit < 0 {
		return 0
	}
	longest := 1
	for j := 0; j < len(nums); {
		length := 1
		max, min := nums[j], nums[j]
		maxI, minI := j, j
		i := 0
		for i = j + 1; i < len(nums); i ++ {
			if limit >= max - nums[i] && limit >= nums[i] - min {
				length += 1
				if max <= nums[i] {
					max = nums[i]
					maxI = i
				}
				if min >= nums[i] {
					min = nums[i]
					minI = i
				}
			} else {
				if longest < length {
					longest = length
				}
				if max <= nums[i] {
					j = minI + 1
				} else if min > nums[i] {
					j = maxI + 1
				}
				break
			}
		}
		if i == len(nums) {
			if longest < length {
				longest = length
			}
			break
		}

	}

	return longest
}
