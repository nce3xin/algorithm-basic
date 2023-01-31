package main

import "fmt"

const N int = 1e5 + 10

var n, m int
var q, d [N]int        // d[]: 每个点的入度，q: BFS队列
var hh, tt int = 0, -1 // 手动维护队列结构需要的变量
var h, e, ne [N]int    // 保存有向图的结构
var idx int = 0

func main() {
	fmt.Scanf("%d%d", &n, &m)
	// init h[]
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	for ; m > 0; m-- {
		var a, b int
		fmt.Scanf("%d%d", &a, &b)
		add(a, b)
		d[b]++ // 添加一条从a指向b的点，b的入度+1
	}
	if topsort() {
		// 如果有拓扑排序，队列里的值就是拓扑排序的值
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", q[i])
		}
	} else {
		fmt.Printf("-1")
	}
}

func topsort() bool {
	for i := 1; i <= n; i++ {
		if d[i] == 0 { // 入度为0的点可以作为起点，入队
			tt++
			q[tt] = i
		}
	}
	for hh <= tt { // 队列不为空
		// 取出队头
		t := q[hh]
		hh++
		// 枚举队头的所有出边
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			d[j]--
			if d[j] == 0 {
				tt++
				q[tt] = j
			}
		}
	}
	return tt == n-1
}

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}
