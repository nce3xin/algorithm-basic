package main

import "container/heap"

type maxHeap []int
type minHeap []int

// 对顶堆
var down *maxHeap = &maxHeap{} // 下方的大根堆，保留较小的那部分数
var up *minHeap = &minHeap{}   // 上方的小根堆，保留较大的那部分数

// 每次数据流新读入一个数
func insert(num int) {
	if down.Empty() || num <= down.Top().(int) {
		heap.Push(down, num)
	} else {
		heap.Push(up, num)
	}
	// 插入新数之后，调整堆，使两个堆里的元素个数相等，或者下方比上方的堆多1
	if down.Len() > up.Len()+1 {
		heap.Push(up, down.Top())
		heap.Pop(down)
	}
	if up.Len() > down.Len() {
		heap.Push(down, up.Top())
		heap.Pop(up)
	}
}

// 求中位数
func getMedian() float64 {
	s := down.Len() + up.Len()
	if s%2 == 0 {
		return float64(down.Top().(int)+up.Top().(int)) / 2
	}
	return float64(down.Top().(int))
}

// max heap

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h *maxHeap) Top() interface{} {
	return (*h)[0]
}

func (h *maxHeap) Empty() bool {
	return h.Len() == 0
}

// min heap

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h *minHeap) Top() interface{} {
	return (*h)[0]
}

func (h *minHeap) Empty() bool {
	return h.Len() == 0
}
