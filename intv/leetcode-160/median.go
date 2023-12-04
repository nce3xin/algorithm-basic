package main

import "container/heap"

type maxHeap []int
type minHeap []int

type MedianFinder struct {
	down *maxHeap // 下方的大根堆，保留较小的那部分数
	up   *minHeap // 上方的小根堆，保留较大的那部分数
}

/** initialize your data structure here. */

func Constructor() MedianFinder {
	return MedianFinder{
		down: &maxHeap{},
		up:   &minHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.down.Empty() || num <= this.down.Top().(int) {
		heap.Push(this.down, num)
	} else {
		heap.Push(this.up, num)
	}
	// 插入新数之后，调整堆，使两个堆里的元素个数相等，或者下方比上方的堆多1
	if this.down.Len() > this.up.Len()+1 {
		heap.Push(this.up, this.down.Top())
		heap.Pop(this.down)
	}
	if this.up.Len() > this.down.Len() {
		heap.Push(this.down, this.up.Top())
		heap.Pop(this.up)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	s := this.down.Len() + this.up.Len()
	if s%2 == 0 {
		return float64(this.down.Top().(int)+this.up.Top().(int)) / 2
	}
	return float64(this.down.Top().(int))
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

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
