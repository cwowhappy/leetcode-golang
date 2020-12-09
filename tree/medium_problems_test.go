package tree

import "testing"

func TestInorderTraversal(t *testing.T) {
	var root, lChild, rChild *TreeNode
	lChild = &TreeNode{Val: 1}
	rChild = &TreeNode{Val: 3}
	root = &TreeNode{Val: 2, Left: lChild, Right: rChild}
	caseResult := []int{1, 2, 3}
	result := InorderTraversalV2(root)
	if len(result) == len(caseResult) {
		for i := 0; i < 3; i ++ {
			if result[i] != caseResult[i] {
				t.Errorf("%d不一致<%d, %d>", i, result[i], caseResult[i])
				break
			}
		}
	} else {
		t.Errorf("长度不一致 result length = %d result = %v", len(result), result)
	}
}

func TestPreorderTraversal(t *testing.T) {
	var root, lChild, rChild *TreeNode
	lChild = &TreeNode{Val: 1}
	rChild = &TreeNode{Val: 3}
	root = &TreeNode{Val: 2, Left: lChild, Right: rChild}
	caseResult := []int{2, 1, 3}
	result := PreorderTraversalV2(root)
	if len(result) == len(caseResult) {
		for i := 0; i < 3; i ++ {
			if result[i] != caseResult[i] {
				t.Errorf("%d不一致<%d, %d>", i, result[i], caseResult[i])
				break
			}
		}
	} else {
		t.Errorf("长度不一致 result length = %d result = %v", len(result), result)
	}
}
