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

/**
Problem849 到最近的人的距离
 */
func MaxDistToClosest(seats []int) int {
	tmpDistance := 0
	maxDistance := 0
	prePeoPleIndex := -1
	for i := 0; i < len(seats); i ++ {
		if seats[i] == 0 {
			continue
		}
		if prePeoPleIndex >= 0 {
			tmpDistance = (i - prePeoPleIndex) / 2
			if maxDistance < tmpDistance {
				maxDistance = tmpDistance
			}
		} else {
			maxDistance = i
		}
		prePeoPleIndex = i
	}
	tmpDistance = len(seats) - 1 - prePeoPleIndex
	if maxDistance < tmpDistance {
		maxDistance = tmpDistance
	}

	return maxDistance
}

/**
Problem954 二倍数对数组
 */
func CanReorderDoubled(A []int) bool {
	can := true
	var needExchange bool
	aLength := len(A)
	for i := 0; i < aLength; i += 2 {
		a := A[i]
		needExchange = false
		j := i + 1
		for j < aLength {
			if a * 2 == A[j] {
				break
			}
			if a == A[j] * 2 {
				needExchange = true
				break
			}
		}
		if j == aLength {
			can = false
			break
		} else {
			A[i + 1], A[j] = A[j], A[i + 1]
			if needExchange {
				A[i + 1], A[i] = A[i], A[i + 1]
			}
		}
	}
	return can
}


