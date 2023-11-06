package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 存每个值在中序遍历中的下标
var hash = make(map[int]int, 0)

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	// init hash
	for i := 0; i < n; i++ {
		hash[inorder[i]] = i
	}
	return dfs(preorder, inorder, 0, n-1, 0, n-1)
}

// pl: preorder left
// pr: preorder right
// il: inorder left
// ir: inorder right
func dfs(preorder []int, inorder []int, pl, pr, il, ir int) *TreeNode {
	if pl > pr {
		return nil
	}
	// k: 根节点在中序遍历中的下标减去il，也是是左子树中序遍历的长度
	k := hash[preorder[pl]] - il
	left := dfs(preorder, inorder, pl+1, pl+k, il, il+k-1)
	right := dfs(preorder, inorder, pl+k+1, pr, il+k+1, ir)
	root := NewTreeNode(preorder[pl], left, right)
	return root
}

func NewTreeNode(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}
