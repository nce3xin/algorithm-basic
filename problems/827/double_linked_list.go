package main

import "fmt"

const N int = 1e5 + 10

var e []int = make([]int, N) // e[i]表示节点i的值
var l []int = make([]int, N) // l[i]表示节点i的前一个点的下标
var r []int = make([]int, N) // r[i]表示节点i的后一个点的下标
var idx int = 2              // 因为0用作head，1用作tail，所以idx从2开始

func init() {
	// 下标为0的是head，下标为1的是tail
	// 和单链表不一样，单链表需要一个head头指针，保存头结点的下标
	// 但双链表的head和tail不保存元素，不算是双链表内的元素
	r[0] = 1
	l[1] = 0
}

func main() {
	var M int
	fmt.Scanf("%d", &M)
	var op string
	for ; M > 0; M-- {
		fmt.Scanf("%s", &op)
		if op == "L" {
			var x int
			fmt.Scanf("%d", &x)
			insertRight(0, x)
		} else if op == "R" {
			var x int
			fmt.Scanf("%d", &x)
			insertLeft(1, x)
		} else if op == "D" {
			var k int
			fmt.Scanf("%d", &k)
			delete(k + 1)
		} else if op == "IL" {
			var k, x int
			fmt.Scanf("%d%d", &k, &x)
			insertLeft(k+1, x)
		} else if op == "IR" {
			var k, x int
			fmt.Scanf("%d%d", &k, &x)
			insertRight(k+1, x)
		}
	}

	for i := r[0]; i != 1; i = r[i] {
		fmt.Printf("%d ", e[i])
	}
}

// 在下标是k的点的右边，插入一个点
func insertRight(k, x int) {
	e[idx] = x
	r[idx] = r[k]
	l[idx] = k
	l[r[k]] = idx
	r[k] = idx
	idx++
}

// 在下标是k的点的左边，插入一个点
func insertLeft(k, x int) {
	insertRight(l[k], x)
}

// 删除下标为k的点
func delete(k int) {
	l[r[k]] = l[k]
	r[l[k]] = r[k]
}
