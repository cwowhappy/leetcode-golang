package array

/**
Problem867 矩阵转换
 */
func transpose(A [][]int) [][]int {
	aLength := len(A)
	bLength := len(A[0])
	resultA := make([][]int, bLength)

	for i := 0; i < bLength; i ++ {
		resultA[i] = make([]int, aLength)
		for j := 0; j < aLength; j ++ {
			resultA[i][j] = A[j][i]
		}
	}

	return resultA
}

/**
Problem977 有序数组的平方
 */
func sortedSquares(nums []int) []int {
	var numsLength = len(nums)
	var i, j int
	k := 0
	resultNums := make([]int, numsLength)
	for i = 0; i < numsLength; i ++ {
		if nums[i] >= 0 {
			break
		}
	}
	j = i - 1
	for j >= 0 && i < numsLength {
		if -nums[j] > nums[i] {
			resultNums[k] = nums[i] * nums[i]
			i += 1
		} else {
			resultNums[k] = nums[j] * nums[j]
			j -= 1
		}
		k += 1
	}

	for j >= 0 {
		resultNums[k] = nums[j] * nums[j]
		j -= 1
		k += 1
	}

	for i < numsLength {
		resultNums[k] = nums[i] * nums[i]
		i += 1
		k += 1
	}

	return resultNums
}

/**
面试题17.10 主要元素
 */
func majorityElementSimple(nums []int) int {
	//直接用map实现
	var mainNum = -1
	numsLength := len(nums)
	numMap := make(map[int]int)
	for i := 0; i < numsLength; i ++ {
		freq, _ := numMap[nums[i]]
		numMap[nums[i]] = freq + 1
		if numMap[nums[i]] * 2 > numsLength {
			mainNum = nums[i]
			break
		}
	}
	return mainNum
}

/**
面试题17.10 主要元素
附加条件：时间复杂度为O(N)，空间复杂度为O(1)
*/
func MajorityElementHard(nums []int) int {
	numsLength := len(nums)
	mainNum := -1
	cnt := 0
	for i := 0; i < numsLength; i ++ {
		if cnt == 0 {
			mainNum = nums[i]
			cnt = 1
		} else {
			if mainNum == nums[i] {
				cnt ++
			} else {
				cnt --
			}
		}
	}
	if cnt == 0 {
		mainNum = -1
	} else {
		cnt = 0
		for i := 0; i < numsLength; i ++ {
			if mainNum == nums[i] {
				cnt ++
			}
		}
		if cnt * 2 < numsLength {
			mainNum = -1
		}
	}
	return mainNum
}
