package tree

/**
Problem124
 */
func MaxPathSum(root *TreeNode) int {
	_, maxPathSum := recPathSum(root)
	return maxPathSum
}

func recPathSum(root *TreeNode) (pathSum int, maxPathSum int) {
	if nil == root {
		pathSum = 0
		maxPathSum =  ^int(^uint(0) >> 1)
		return
	}
	lPathSum, lMaxPathSum := recPathSum(root.Left)
	rPathSum, rMaxPathSum := recPathSum(root.Right)
	sum := maxInt(0, lPathSum) + maxInt(0, rPathSum) + root.Val
	maxPathSum = maxInt(lMaxPathSum, rMaxPathSum, sum)
	pathSum = maxInt(0, lPathSum, rPathSum) + root.Val
	return
}

func maxInt(nums ...int) int{
	//asset len(nums) > 0
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxNum < nums[i] {
			maxNum = nums[i]
		}
	}
	return maxNum
}

