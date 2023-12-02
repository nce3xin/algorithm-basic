package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthNode(root *TreeNode, k int) *TreeNode {
	var res *TreeNode
	dfs(root, &k, &res)
	return res
}

func dfs(root *TreeNode, k *int, res **TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left, k, res)
	// 中序遍历，举个例子，这里每次遍历到的节点为 [1,2,3,4,5,6]
	// 第一次走到这个位置，遍历的是1。第二次走到这个位置，遍历的是2
	*k = (*k) - 1
	if *k == 0 {
		*res = root
	}
	dfs(root.Right, k, res)
}
