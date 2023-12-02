package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	// 先判断pRoot2是否是以pRoot1为根的子树
	if isPart(pRoot1, pRoot2) {
		return true
	}
	// 如果不是，递归判断pRoot2是否是以pRoot1的左右子树为根的子树
	// 这儿递归的是hasSubtree(), 不是isPart()
	return hasSubtree(pRoot1.Left, pRoot2) || hasSubtree(pRoot1.Right, pRoot2)
}

// p2是否是以p1为根的子树
func isPart(p1, p2 *TreeNode) bool {
	if p2 == nil {
		return true
	}
	// p1为空，但p2不为空（由第一个if语句保证）
	if p1 == nil {
		return false
	}
	if p1.Val != p2.Val {
		return false
	}
	// 如果p1和p2相等，递归判断p1,p2的左右子树是否相等
	return isPart(p1.Left, p2.Left) && isPart(p1.Right, p2.Right)
}
