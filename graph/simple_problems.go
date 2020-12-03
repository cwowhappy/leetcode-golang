package graph

/**
Problem997 找到小镇的法官
时间复杂度: O(n)
空间复杂度: O(n)
 */
func findJudge(N int, trusts [][]int) int {
	inDegrees := make([]int, N)
	outDegrees := make([]int, N)
	for i := 0; i < len(trusts); i ++ {
		inDegrees[trusts[i][1]] ++
		outDegrees[trusts[i][0]] ++
	}
	index := -1
	for i := 0; i < N; i ++ {
		if outDegrees[i] == 0 && inDegrees[i] == N - 1 {
			index = i
			break
		}
	}
	return index
}