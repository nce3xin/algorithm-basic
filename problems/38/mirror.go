package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirror(root *TreeNode) {
	if root == nil {
		return
	}
	mirror(root.Left)
	mirror(root.Right)
	// swap left and right
	root.Left, root.Right = root.Right, root.Left
}
