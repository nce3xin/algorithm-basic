package main

import "fmt"

const N int = 1e5 + 10

var n, m int
var h, e, ne [N]int
var idx int = 0
var q [N]int
var d [N]int // 每个点到原点的距离

func main() {
	// init h[] and d[]
	for i := 0; i < N; i++ {
		h[i] = -1
		d[i] = -1
	}

	d[1] = 0 //存储每个节点离起点的距离

	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanf("%d%d", &a, &b)
		add(a, b)
	}
	fmt.Printf("%d", bfs())
}

func bfs() int {
	hh, tt := 0, -1
	// Push first one into queue
	tt++
	q[tt] = 1      //0号节点是编号为1的节点
	for hh <= tt { // queue is not empty
		//取出队列头部节点
		t := q[hh]
		hh++
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if d[j] == -1 { // j没有被遍历过
				d[j] = d[t] + 1
				tt++
				q[tt] = j
			}
		}
	}
	return d[n]
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}
