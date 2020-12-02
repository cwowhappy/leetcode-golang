package array

//Problem1523
func GetWinner(arr []int, k int) int {
	winnerNum := arr[0]
	winEpoch := 0
	for i := 1; i < len(arr) && winEpoch < k; i ++ {
		if winnerNum >= arr[i] {
			winEpoch ++
		} else {
			winnerNum = arr[i]
			winEpoch = 1
		}
	}
	return winnerNum
}

/**
Problem1552 两球之间的磁力
 */
func MaxDistance(position []int, m int) int {
	return 0
}

