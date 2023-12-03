package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findPath(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	dfs(root, sum, &res, &path)
	return res
}

func dfs(root *TreeNode, sum int, res *[][]int, path *[]int) {
	if root == nil {
		return
	}
	// 这里做减法，不做加法，可以少传一个参数
	sum -= root.Val
	*path = append(*path, root.Val)
	// 如何判断当前节点是叶子节点呢？就是判断左右儿子是否都为空
	if root.Left == nil && root.Right == nil && sum == 0 {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*res = append(*res, tmp)
	}
	dfs(root.Left, sum, res, path)
	dfs(root.Right, sum, res, path)
	// 恢复现场，去掉刚push进path的元素
	// sum不用恢复，因为它是个标量（局部变量），每个dfs都有自己的sum
	*path = (*path)[:len(*path)-1]
}
