package linkedlist

const N int = 1e5 + 10

var e []int = make([]int, N)  // ne[i]表示节点i的next指针是多少
var ne []int = make([]int, N) // e[i]表示节点i的值
var idx int = 0               // 存储当前已经用到了哪个下标
var head int = -1             // head表示头结点的下标，初始化为-1

// 向链表头插入一个数
func insertToHead(x int) {
	e[idx] = x
	ne[idx] = head
	head = idx
	idx++
}

// 删除下标为k的后面的点
func delete(k int) {
	ne[k] = ne[ne[k]]
}

// 将x插到下标为k的点后面
func insert(k, x int) {
	e[idx] = x
	ne[idx] = ne[k]
	ne[k] = idx
	idx++
}
