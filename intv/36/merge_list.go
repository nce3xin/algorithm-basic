package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// 定义合并之后的链表的虚拟头节点
	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	// cur记录当前链表的尾节点在哪
	cur := dummy
	// 2个链表都不为空
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}
	// 走到这儿说明，至少有一个链表为空了，把不为空的链表接过来
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
