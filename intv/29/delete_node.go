package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(head *ListNode) *ListNode {
	// 这题可能会把头节点删掉
	// 凡是可能会把头节点删掉的问题，一般都会定义一个虚拟头节点来简化代码
	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummy.Next = head
	// 上一个保留的段的结尾, 因为只保留长度为1的段，所以结尾也就是指向这段里的唯一一个元素
	p := dummy
	// 因为下面会用到p.Next，所以这里必须要保证p.Next不为空
	for p.Next != nil { // p不为空就一直做
		// q: 下一段的第一个节点
		q := p.Next
		// 跳过中间相同元素
		for q != nil && p.Next.Val == q.Val {
			q = q.Next
		}
		// 判断下段长度是不是1
		if p.Next.Next == q {
			p = p.Next
		} else { // 下段长度不是1，把下段删掉
			p.Next = q
		}
	}
	return dummy.Next
}
