package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func findFirstCommonNode(headA *ListNode, headB *ListNode) *ListNode {
	p := headA
	q := headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}

		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}