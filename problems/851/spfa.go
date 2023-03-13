package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 10 // 点的数量

var n, m int
var h, e, w, ne [N]int // 稀疏图，用邻接表存储, w存储边的权重
var idx int = 0
var st [N]bool  // 当前的点是否在队列当中。防止队列中存储重复的点
var dist [N]int // 存储所有点到1号点的距离

// queue
var q [N]int
var hh, tt int = 0, -1

func main() {
	// init h[] and dist[]
	for i := 0; i < N; i++ {
		h[i] = -1
		dist[i] = 0x3f3f3f3f // infinity
	}
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	for ; m > 0; m-- {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		add(a, b, c)
	}
	t := spfa()
	if t == 0x3f3f3f3f {
		fmt.Printf("impossible\n")
	} else {
		fmt.Printf("%d\n", t)
	}
}

func add(a, b, c int) {
	w[idx] = c
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func spfa() int {
	dist[1] = 0
	push(1)
	st[1] = true
	for !empty() {
		t := front()
		pop()
		// 点从队列里出来了，不在队列里了，所以st[t] = false
		st[t] = false
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]
				if !st[j] {
					push(j)
					st[j] = true
				}
			}
		}
	}
	if dist[n] == 0x3f3f3f3f {
		return 0x3f3f3f3f
	}
	return dist[n]
}

func push(x int) {
	tt++
	q[tt] = x
}

func pop() {
	hh++
}

func front() int {
	return q[hh]
}

func empty() bool {
	if hh <= tt {
		return false
	}
	return true
}
