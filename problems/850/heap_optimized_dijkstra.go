package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N int = 1e6 + 10 // 点的数量

var n, m int
var h, e, w, ne [N]int // 稀疏图，用邻接表存储, w存储边的权重
var idx int = 0
var st [N]bool  // 记录每个点是否已经找到最短路径
var dist [N]int // 存储所有点到1号点的距离
type PII struct {
	first  int // distance to origin
	second int // number of node
}
type Heap []PII // heap

// 用fmt.Scanf()会超时

func main() {
	// init h[] and dist[]
	for i := 0; i < N; i++ {
		h[i] = -1
		dist[i] = 0x3f3f3f3f // infinity
	}
	fmt.Scanf("%d%d", &n, &m)
	r := bufio.NewReader(os.Stdin)
	for ; m > 0; m-- {
		in := readline(r)
		//fmt.Scanf("%d%d%d", &x, &y, &z)
		x, y, z := in[0], in[1], in[2]
		add(x, y, z)
	}
	d := dijkstra()
	fmt.Printf("%d", d)
}

func dijkstra() int {
	pq := Heap{}              // heap, aka priority queue
	heap.Push(&pq, PII{0, 1}) // first存储距离，second存储节点编号
	dist[1] = 0
	for pq.Len() > 0 { // pq is not empty
		x := heap.Pop(&pq).(PII)
		ver := x.second // vertex
		distance := x.first
		if st[ver] { // 如果j点已经找到最短路径了，就跳过
			continue
		}
		st[ver] = true // 标记
		for i := h[ver]; i != -1; i = ne[i] {
			j := e[i] // i只是个下标，e中在存的是i这个下标对应的点。
			if dist[j] > distance+w[i] {
				dist[j] = distance + w[i]
				heap.Push(&pq, PII{dist[j], j})
			}
		}
	}
	if dist[n] == 0x3f3f3f3f {
		return -1
	}
	return dist[n]
}

func add(a, b, c int) {
	w[idx] = c
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(PII))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	// 小根堆
	return (*h)[i].first < (*h)[j].first
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func readline(r *bufio.Reader) []int {
	s, _ := r.ReadString('\n')
	ss := strings.Fields(s)
	res := make([]int, len(ss))
	for i, v := range ss {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}
