package main

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Father *TreeNode
}

func inorderSuccessor(p *TreeNode) *TreeNode {
	// 如果右子树存在，右子树最左边的点就是后继
	if p.Right != nil {
		p = p.Right // 移动到右子树上
		// 在右子树上一直往左走
		for p.Left != nil {
			p = p.Left
		}
		return p
	}
	// 如果没有右子树，需要一直往左上方走
	for p.Father != nil && p == p.Father.Right { // p是father的右儿子的话，一直往左上走
		p = p.Father
	}
	return p.Father
}
