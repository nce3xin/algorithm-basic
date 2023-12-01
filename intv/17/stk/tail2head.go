package main

type ListNode struct {
	Val  int
	Next *ListNode
}

const N int = 1010

func printListReversingly(head *ListNode) []int {
	var stk [N]int
	tt := 0
	for cur := head; cur != nil; cur = cur.Next {
		tt++
		stk[tt] = cur.Val
	}
	res := make([]int, 0)
	for tt > 0 {
		x := stk[tt]
		tt--
		res = append(res, x)
	}
	return res
}
