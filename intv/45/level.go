package main

const N int = 1010

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printFromTopToBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := [N]*TreeNode{}
	hh, tt := 0, 0
	q[0] = root
	zigzag := false // false表示从左往右打印，true是从右往左打印
	for hh <= tt {
		size := tt - hh + 1
		tmp := make([]int, 0)
		for i := 0; i < size; i++ {
			t := q[hh]
			hh++
			tmp = append(tmp, t.Val)
			if t.Left != nil {
				tt++
				q[tt] = t.Left
			}
			if t.Right != nil {
				tt++
				q[tt] = t.Right
			}
		}
		if zigzag {
			reverse(tmp)
		}
		zigzag = !zigzag
		res = append(res, tmp)
	}
	return res
}

func reverse(a []int) {
	n := len(a)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
