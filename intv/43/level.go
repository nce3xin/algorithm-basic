package main

const N int = 1010

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printFromTopToBottom(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := [N]*TreeNode{}
	hh, tt := 0, 0
	q[0] = root
	for hh <= tt {
		t := q[hh]
		hh++
		res = append(res, t.Val)
		if t.Left != nil {
			tt++
			q[tt] = t.Left
		}
		if t.Right != nil {
			tt++
			q[tt] = t.Right
		}
	}
	return res
}
