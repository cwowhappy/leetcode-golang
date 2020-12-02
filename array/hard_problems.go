package array

/**
面试题17.21
思路一：
时间复杂度:O(n)
空间复杂度:O(n)
*/
func Trap(heights []int) int {
	totalValue := 0
	value := 0
	values := make([]int, 0, len(heights))
	maxHeight := 0
	for i := 0; i < len(heights); i ++ {
		if value = maxHeight - heights[i]; value > 0 {
			values = append(values, value)
		} else {
			maxHeight = heights[i]
			values = append(values, 0)
		}
	}

	maxHeight = 0
	for i := len(heights) - 1; i >= 0; i -- {
		if value = maxHeight - heights[i]; value > 0 {
			if value < values[i] {
				values[i] = value
				totalValue += value
			} else {
				totalValue += values[i]
			}
		} else {
			maxHeight = heights[i]
		}
	}
	return totalValue
}

/**
面试题17.21
思路二:
时间复杂度:O(n)
空间复杂度:O(1)
*/
func TrapV2(heights []int) int {
	totalValue := 0
	heightsLength := len(heights)
	if heightsLength > 2 {
		maxValueL, maxValueR := heights[0], heights[heightsLength - 1]
		for i, j := 1, heightsLength - 2; i <= j; {
			if maxValueL < maxValueR {
				if maxValueL < heights[i] {
					maxValueL = heights[i]
				} else {
					totalValue += maxValueL - heights[i]
				}
				i ++
			} else {
				if maxValueR < heights[j] {
					maxValueR = heights[j]
				} else {
					totalValue += maxValueR - heights[j]
				}
				j --
			}
		}
	}
	return totalValue
}
