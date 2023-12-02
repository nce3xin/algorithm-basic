package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(treeDepth(root.Left)+1, treeDepth(root.Right)+1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
