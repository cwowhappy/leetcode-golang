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