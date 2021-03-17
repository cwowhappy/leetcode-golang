package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := 4
	if result != LengthOfLIS(nums) {
		t.Error(nums, result)
	}
}

func TestFindCheapestPrice(t *testing.T) {
	result := 200
	n, src, dst, k := 3, 0, 2, 1
	flights := [][]int{
		{0, 1, 100}, {1, 2, 100}, {0, 2, 500},
	}
	if result != FindCheapestPrice(n, flights, src, dst, k) {
		t.Error("error")
	}
}

func TestMinHeightShelves(t *testing.T) {
	books := [][]int{{1,1},{2,3},{2,3},{1,1},{1,1},{1,1},{1,2}}
	shelfWidth := 4
	result := 6
	if result != MinHeightShelves(books, shelfWidth) {

	}
}

func TestMatrixBlockSum(t *testing.T) {
	K := 1
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	t.Error(MatrixBlockSum(mat, K))
}

func TestFindTargetSumWays(t *testing.T) {
	S := 3
	nums := []int{1, 1, 1, 1, 1}
	result := 5
	realResult := FindTargetSumWays(nums, S)
	if result != realResult {
		t.Error(nums, S, result, realResult)
	}
}

func TestCuttingRope(t *testing.T) {
	var n, expectResult int
	//n = 2
	//expectResult = 1
	//if expectResult != CuttingRope(n) {
	//	t.Error()
	//}

	n = 120
	expectResult = 953271190
	if expectResult != CuttingRope(n) {
		t.Error()
	}

}