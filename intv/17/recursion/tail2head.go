package main

type ListNode struct {
	Val  int
	Next *ListNode
}

const N int = 1010

func printListReversingly(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	if head.Next == nil {
		return []int{head.Val}
	}
	t := printListReversingly(head.Next)
	t = append(t, head.Val)
	return t
}
