package unclassified

/**
Problem55 跳跃游戏
时间复杂度:
空间复杂度:
 */
func CanJump(nums []int) bool {
	maxJumpIndex := 0
	for i := 0; i <= maxJumpIndex && maxJumpIndex < len(nums); i++ {
		if nums[i] + i > maxJumpIndex {
			maxJumpIndex = nums[i] + i
		}
	}
	return maxJumpIndex >= len(nums) - 1
}

/**
Problem45 跳跃游戏II
 */
func Jump(nums []int) int {
	numsLength := len(nums)
	step := 0
	maxLength := 0
	indexBoundaries := [][]int{{0, 0}, {0, 0}}
	for maxLength < numsLength - 1 {
		indexBoundaryI := step & 1
		indexBoundaries[indexBoundaryI^1][0] = indexBoundaries[indexBoundaryI][1] + 1
		for i := indexBoundaries[indexBoundaryI][0]; i <= indexBoundaries[indexBoundaryI][1] && i < numsLength; i ++ {
			if i + nums[i] > indexBoundaries[indexBoundaryI^1][1] {
				indexBoundaries[indexBoundaryI^1][1] = i + nums[i]
			}
		}
		maxLength = indexBoundaries[indexBoundaryI^1][1]
		step ++
	}
	return step
}

func JumpV2(nums []int) int {
	step := 0
	end := 0
	maxIndex := 0
	for i := 0; i < len(nums) - 1 && end < len(nums) - 1; i ++ {
		if i + nums[i] > maxIndex {
			maxIndex = i + nums[i]
		}
		if i == end {
			end = maxIndex
			step ++
		}
	}

	return step
}

/**
Problem 22
 */
func GenerateParenthesis(n int) []string {
	return nil
}