package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 判断左右子树是否互为镜像
	return dfs(root.Left, root.Right)
}

func dfs(p, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == nil && q == nil {
			return true
		}
		return false
	}
	return p.Val == q.Val && dfs(p.Left, q.Right) && dfs(p.Right, q.Left)
}
