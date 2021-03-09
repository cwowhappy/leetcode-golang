package dp

import "fmt"

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

/**
Problem787
 */
func FindCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	flightsLength := len(flights)
	distMatrix := make([][]int, 2)
	distMatrix[0] = make([]int, n)
	distMatrix[1] = make([]int, n)
	for i := 0; i < n; i ++ {
		distMatrix[0][i] = -1
		distMatrix[1][i] = -1
	}

	distMatrix[0][src], distMatrix[1][src] = 0, 0
	for k := 0; k <= K; k ++ {
		index := k & 1
		for i := 0; i < flightsLength; i ++ {
			if distMatrix[index^1][flights[i][0]] != -1 {
				if distMatrix[index][flights[i][1]] == -1 || distMatrix[index][flights[i][1]] > distMatrix[index^1][flights[i][0]] + flights[i][2] {
					distMatrix[index][flights[i][1]] = distMatrix[index^1][flights[i][0]] + flights[i][2]
				}
			}
		}
		fmt.Printf("========%d=======\n", k)
		fmt.Println(distMatrix[0])
		fmt.Println(distMatrix[1])
	}

	return distMatrix[K&1][dst]
}

/**
Problem 1105
 */
func MinHeightShelves(books [][]int, shelfWidth int) int {
	bookLen := len(books)
	minHeights := make([]int, 0, bookLen + 1)
	minHeights = append(minHeights, 0)
	for i := 1; i <= bookLen; i ++ {
		minHeights = append(minHeights, 1000 * 1000)
	}

	for i := 1; i <= bookLen; i++ {
		maxH := books[i - 1][1]
		for j, tmpWidth := i, books[i - 1][0]; j > 0 && shelfWidth >= tmpWidth; {
			minHeights[i] = min2(minHeights[i], minHeights[j - 1] + maxH)
			j -= 1
			if j == 0 {
				break
			}
			maxH = max2(maxH, books[j - 1][1])
			tmpWidth += books[j - 1][0]
		}
	}
	return minHeights[bookLen]
}

/**
Problem 1314
 */
func MatrixBlockSum(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
	tmpSums := make([]int, 0, n)
	sumMatrix := make([][]int, 0, m)
	for j := 0; j < m; j ++ {
		sumMatrix = append(sumMatrix, make([]int, 0, n))
		for i := 0; i < n; i ++ {
			sumMatrix[j] = append(sumMatrix[j], 0)
		}
	}
	for i := 0; i < n; i ++ {
		tmpSums = append(tmpSums, mat[0][i])
	}
	for j := 1; j < K && j < m; j ++ {
		for i := 0; i < n; i ++ {
			tmpSums[i] += mat[j][i]
		}
	}

	for j := 0; j < m; j ++ {
		if j - K - 1 >= 0 {
			for i := 0; i < n; i ++ {
				tmpSums[i] -= mat[j - K - 1][i]
			}
		}
		if j + K < m {
			for i := 0; i < n; i ++ {
				tmpSums[i] += mat[j + K][i]
			}
		}
		tmpTotalSum := 0
		for k := 0; k < K  && k < n; k ++ {
			tmpTotalSum += tmpSums[k]
		}
		for i := 0; i < n; i ++ {
			if i + K < n {
				tmpTotalSum += tmpSums[i+K]
			}
			if i - K - 1 >= 0 {
				tmpTotalSum -= tmpSums[i-K-1]
			}
			sumMatrix[j][i] = tmpTotalSum
		}
	}
	return sumMatrix
}

/**
Problem 494
 */
func FindTargetSumWays(nums []int, S int) int {
	var targetV int
	var ok bool
	sumMaps := make([]map[int]int, 0, 2)
	sumMaps = append(sumMaps, make(map[int]int))
	sumMaps = append(sumMaps, make(map[int]int))

	if nums[0] == 0 {
		sumMaps[0][nums[0]] = 2
	} else {
		sumMaps[0][nums[0]] = 1
		sumMaps[0][-nums[0]] = 1
	}
	source, target := 1, 0
	for i := 1; i < len(nums); i ++ {
		source, target = target, source
		for k, v := range sumMaps[source] {
			if targetV, ok = sumMaps[target][k + nums[i]]; !ok {
				targetV = 0
			}
			sumMaps[target][k + nums[i]] = targetV + v
			if targetV, ok = sumMaps[target][k - nums[i]]; !ok {
				targetV = 0
			}
			sumMaps[target][k - nums[i]] = targetV + v
		}
	}
	if targetV, ok = sumMaps[target][S]; !ok {
		targetV = 0
	}
	return targetV
}