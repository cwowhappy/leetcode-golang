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
