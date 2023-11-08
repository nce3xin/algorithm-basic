package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func findKthToTail(pListHead *ListNode, k int) *ListNode {
	n := 0
	for p := pListHead; p != nil; p = p.Next {
		n++
	}
	if k > n {
		return nil
	}
	p := pListHead
	for i := 0; i < n-k; i++ {
		p = p.Next
	}
	return p
}
