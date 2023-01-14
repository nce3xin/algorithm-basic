package main

import "fmt"

const N int = 1e5 + 10

var e []int = make([]int, N)  // ne[i]表示节点i的next指针是多少
var ne []int = make([]int, N) // e[i]表示节点i的值
var idx int = 0               // 存储当前已经用到了哪个下标
var head int = -1             // head表示头结点的下标，初始化为-1

func main() {
	var M int
	fmt.Scanf("%d", &M)
	var op string
	for ; M > 0; M-- {
		fmt.Scanf("%s", &op)
		if op == "H" {
			var x int
			fmt.Scanf("%d", &x)
			insertToHead(x)
		} else if op == "I" {
			var k, x int
			fmt.Scanf("%d", &k)
			fmt.Scanf("%d", &x)
			insert(k-1, x)
		} else if op == "D" {
			var k int
			fmt.Scanf("%d", &k)
			// k == 0 表示删除头节点，直接让head指向下一个节点
			if k == 0 {
				head = ne[head]
			} else {
				delete(k - 1)
			}
		}
	}

	for i := head; i != -1; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}

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
