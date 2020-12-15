package dp

/**
Problem70 爬楼梯
 */
func ClimbStairs(n int) int {
	if n == 0 {
		return 0
	}
	n1, n2 := 1, 0
	for i := 1; i < n; i ++ {
		n2, n1 = n1, n1 + n2
	}
	return n1 + n2
}


/**
Problem53
*/
func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	preSubArraySum := nums[0]
	for i := 1; i < len(nums); i ++ {
		if preSubArraySum <= 0 {
			preSubArraySum = nums[i]
		} else {
			preSubArraySum += nums[i]
		}
		if preSubArraySum > maxSum {
			maxSum = preSubArraySum
		}
	}
	return maxSum
}

/**
Pblem392
 */
func IsSubsequence(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)
	indexS := 0
	for indexT := 0; indexT < lenT; indexT ++ {
		if indexS == lenS {
			break
		} else if s[indexS] == t[indexT] {
			indexS ++
		}
	}

	return indexS == lenS
}

/**
Problem746
 */
func MinCostClimbingStairs(costs []int) int {
	costsLen := len(costs)
	if costsLen <= 1 {
		return 0
	}
	minCostWithCurStep := costs[0]
	minCostWithoutCurStep := 0
	for i := 1; i < costsLen; i ++ {
		tmp := minCostWithCurStep
		minCostWithCurStep = min2(minCostWithoutCurStep, minCostWithCurStep) + costs[i]
		minCostWithoutCurStep = tmp
	}
	return min2(minCostWithCurStep, minCostWithoutCurStep)
}


/**
Problem198
 */
func Rob(nums []int) int {
	numsLen := len(nums)
	if numsLen < 1 {
		return 0
	}
	maxAmountWithCurNode := nums[0]
	maxAmountWithoutCurNode := 0

	for i := 1; i < numsLen; i ++ {
		tmp := maxAmountWithCurNode
		maxAmountWithCurNode = maxAmountWithoutCurNode + nums[i]
		maxAmountWithoutCurNode = max2(tmp, maxAmountWithoutCurNode)
	}
	return max2(maxAmountWithoutCurNode, maxAmountWithCurNode)

}


func min2(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max2(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
