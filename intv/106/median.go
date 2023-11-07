package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type minHeap []int
type maxHeap []int

func main() {
	in := bufio.NewReader(os.Stdin)
	var T int // T为数据集个数
	fmt.Fscan(in, &T)
	for ; T > 0; T-- {
		var n, m int // n表示数据个数，m表示数据集编号
		fmt.Fscan(in, &m, &n)
		// 输出数据集编号和中位数个数(即奇数位个数)
		fmt.Printf("%d %d\n", m, (n+1)/2)
		up := &minHeap{}   // 上方的小根堆
		down := &maxHeap{} // 下方的大根堆
		cnt := 0           // 用于分隔输出,每十个数一行
		for i := 1; i <= n; i++ {
			var x int
			fmt.Fscan(in, &x)
			// 下面为空或x小于等于下方堆顶,则将x插入down
			if down.Empty() || x <= down.Top() {
				heap.Push(down, x)
			} else {
				heap.Push(up, x)
			}

			// 期望：如果有偶数个数，上面和下面一样多，如果有奇数个数，则下面比上面多一个
			if down.Len() > up.Len()+1 { // 下面比上面多了不止一个，把这个放到上面去
				heap.Push(up, down.Top())
				// 注意这儿，不能写成 down.Pop()! 我在这儿卡了一个世纪
				heap.Pop(down)
			}
			if up.Len() > down.Len() { // 上面比下面多一个，把这个放到下面
				heap.Push(down, up.Top())
				// 注意这儿，不能写成 up.Pop()! 我在这儿卡了一个世纪
				heap.Pop(up)
			}

			if i%2 != 0 {
				fmt.Printf("%d ", down.Top()) // 输出down的堆顶
				cnt++
				if cnt%10 == 0 {
					fmt.Println()
				}
			}
		}
		// 如果最后不足10个数的话，下一个数就要另起一行输出
		if cnt%10 != 0 {
			fmt.Printf("\n")
		}
	}
}

func (pq minHeap) Len() int {
	return len(pq)
}

func (pq minHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq minHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *minHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

// Pop 按照接口定义 remove and return element Len() - 1
func (pq *minHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func (pq *minHeap) Empty() bool {
	if pq.Len() == 0 {
		return true
	}
	return false
}

// Top 堆顶元素也就是堆的最小值，保存在队列头
func (pq *minHeap) Top() int {
	return (*pq)[0]
}

// 大根堆

func (pq maxHeap) Len() int {
	return len(pq)
}

func (pq maxHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq maxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *maxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

// Pop 按照接口定义 remove and return element Len() - 1
func (pq *maxHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func (pq *maxHeap) Empty() bool {
	if pq.Len() == 0 {
		return true
	}
	return false
}

// Top 堆顶元素也就是堆的最大值，保存在队列头
func (pq *maxHeap) Top() int {
	return (*pq)[0]
}
