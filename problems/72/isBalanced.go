package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	res := true
	dfs(root, &res)
	return res
}

// 返回值是以root为根的树的最大深度
func dfs(root *TreeNode, res *bool) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)
	if left-right > 1 || right-left > 1 {
		*res = false
	}
	return max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
