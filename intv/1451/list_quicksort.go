package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func quickSortList(head *ListNode) *ListNode {
	// 如果链表里没有数，或者只有1个数
	if head == nil || head.Next == nil {
		return head
	}
	// three dummy heads
	left := NewListNode(-1, nil)
	mid := NewListNode(-1, nil)
	right := NewListNode(-1, nil)
	// record list tail, due to we need to insert item at the tail
	ltail := left
	mtail := mid
	rtail := right
	// pivot
	val := head.Val
	for p := head; p != nil; p = p.Next {
		if p.Val < val {
			ltail.Next = p
			ltail = p
		} else if p.Val == val {
			mtail.Next = p
			mtail = p
		} else {
			rtail.Next = p
			rtail = p
		}
	}
	// point tail's next to nil
	ltail.Next = nil
	mtail.Next = nil
	rtail.Next = nil

	left.Next = quickSortList(left.Next)
	right.Next = quickSortList(right.Next)
	// connect left, mid, right
	// mid并不一定存在，所以我们连接mid和right的时候，仍然从left的头开始找
	// 因为left和mid现在已经连成一个链表了
	// 后来又求证了，mid不会为空，val选择的是链表中的值，在遍历链表的的时候至少会有一个点的值等于val，所以mid链表至少有一个节点
	getTail(left).Next = mid.Next
	getTail(left).Next = right.Next
	return left.Next
}

func getTail(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}
