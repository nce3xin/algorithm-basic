package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

const N int = 2010

// c是辅助数组，存最终答案的，最终必须复制回a数组
// 因为每次是a和b合并，所以必须把每次合并的结果复制回a数组
var a, b, c [N]int
var n, m int

type Item struct {
	s int // 和sum
	i int // 当前元素在a数组里的下标
}

type minHeap []*Item

func main() {
	var T int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &T)
	for ; T > 0; T-- {
		fmt.Fscan(in, &m, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		// 对a数组从小到大排序
		sort.Ints(a[:n])
		// 合并m-1次
		for i := 0; i < m-1; i++ {
			// 每次合并之前先把b数组读入进来
			for j := 0; j < n; j++ {
				fmt.Fscan(in, &b[j]) // 这儿是b[j]不是b[i]，md我卡这儿好久
			}
			// merge()就是把a和b合并，把结果存到a里面
			merge()
		}
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", a[i])
		}
		fmt.Println()
	}
}

func (h minHeap) Less(i, j int) bool {
	return h[i].s < h[j].s
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *minHeap) Top() interface{} {
	return (*h)[0]
}

func merge() {
	h := &minHeap{}
	for i := 0; i < n; i++ {
		heap.Push(h, &Item{
			s: a[0] + b[i],
			i: 0,
		})
	}
	for i := 0; i < n; i++ {
		// 注意这儿不能t:=heap.Pop(h)
		// 因为Pop()出来的数和h.Top()不一样
		t := h.Top().(*Item)
		heap.Pop(h)
		s := t.s
		p := t.i
		c[i] = s
		heap.Push(h, &Item{
			s: s - a[p] + a[p+1],
			i: p + 1,
		})
	}
	for i := 0; i < n; i++ {
		a[i] = c[i]
	}
}
