package unclassified

func Problem001(nums []int) int {
	i, j := 0, 0
	for j < len(nums) {
		tmp := nums[j]
		for k := j; k < len(nums) && k != nums[k]; k = tmp {
			tmp, nums[k] = nums[k], tmp
		}
		if tmp != j {
			nums[j] = -1
			j += 1
		}
		for j < len(nums) && j == nums[j] {
			j += 1
		}
		for i < len(nums) && i == nums[i] {
			i += 1
		}
	}
	return i
}
