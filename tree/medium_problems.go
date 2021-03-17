package tree

import "github.com/cwowhappy/leetcode-golang/util"

/**
Problem110 平衡二叉树
 */
func IsBalanced(root *TreeNode) bool {
	_, flag := recIsBalanced(root)
	return flag
}
func recIsBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	var flag = true
	var height, lHeight, rHeight = 1, 0, 0
	if root.Left != nil {
		lHeight, flag = recIsBalanced(root.Left)
	}
	if flag && root.Right != nil {
		rHeight, flag = recIsBalanced(root.Right)
	}
	if flag {
		if lHeight - rHeight > 1 || rHeight - lHeight > 1 {
			flag = false
		} else {
			if lHeight > rHeight {
				height += lHeight
			} else {
				height += rHeight
			}
		}
	}
	return height, flag
}

/**
Problem144 二叉树的前序遍历
 */
func PreorderTraversal(root *TreeNode) []int {
	var traversalResult *[]int
	tmp := make([]int, 0)
	traversalResult = &tmp
	recPreorderTraversal(root, traversalResult)
	return *traversalResult
}

func recPreorderTraversal(root *TreeNode, traversalResult *[]int) {
	if root == nil {
		return
	}
	*traversalResult = append(*traversalResult, root.Val)
	recPreorderTraversal(root.Left, traversalResult)
	recPreorderTraversal(root.Right, traversalResult)
}

//非递归版:借助queue实现
func PreorderTraversalV2(root *TreeNode) []int {
	traversalResult := make([]int, 0)
	if nil == root {
		return traversalResult
	}
	q := util.Queue{Length: 20}
	_ = q.Enqueue(root)
	for !q.IsEmpty() {
		if queueNode, err := q.Dequeue(); err == nil {
			treeNode := queueNode.(*TreeNode)
			traversalResult = append(traversalResult, treeNode.Val)
			if treeNode.Left != nil {
				_ = q.Enqueue(treeNode.Left)
			}
			if treeNode.Right != nil {
				_ = q.Enqueue(treeNode.Right)
			}
		}
	}

	return traversalResult
}
/**
Problem94 二叉树的中序遍历
 */
func InorderTraversal(root *TreeNode) []int {
	var traversalResult *[]int
	tmp := make([]int, 0)
	traversalResult = &tmp
	recInorderTraversal(root, traversalResult)
	return *traversalResult
}

func recInorderTraversal(root *TreeNode, traversalResult *[]int) {
	if root == nil {
		return
	}
	recInorderTraversal(root.Left, traversalResult)
	*traversalResult = append(*traversalResult, root.Val)
	recInorderTraversal(root.Right, traversalResult)
}

//非递归版:借助stack实现
func InorderTraversalV2(root *TreeNode) []int {
	var traversalResult []int
	s := make(util.Stack, 0)
	for treeNode := root; nil != treeNode; treeNode = treeNode.Left {
		s.Push(treeNode)
	}
	for !s.IsEmpty() {
		if tmp, ok := s.Pop(); ok == nil {
			treeNode := tmp.(*TreeNode)
			traversalResult = append(traversalResult, treeNode.Val)
			for treeNode = treeNode.Right; nil != treeNode; treeNode = treeNode.Left {
				s.Push(treeNode)
			}
		}
	}
	return traversalResult
}

/**
Problem654 最大二叉树
 */
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	return recConstructMaximumBinaryTree(nums, 0, len(nums) - 1)
}

func recConstructMaximumBinaryTree(nums []int, indexBegin int, indexEnd int) *TreeNode {
	if indexBegin > indexEnd {
		return nil
	}
	maxNumIndex := indexBegin
	maxNum := nums[maxNumIndex]
	for index := indexBegin + 1; index <= indexEnd; index ++ {
		if maxNum < nums[index] {
			maxNumIndex = index
			maxNum = nums[maxNumIndex]
		}
	}
	lChild := recConstructMaximumBinaryTree(nums, indexBegin, maxNumIndex - 1)
	rChild := recConstructMaximumBinaryTree(nums, maxNumIndex + 1, indexEnd)

	return &TreeNode{Val: maxNum, Left: lChild, Right: rChild}
}
