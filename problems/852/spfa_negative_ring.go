package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	N int = 1e5 + 10 // 点的数量
)

var n, m int
var h, e, w, ne [N]int // 邻接表存储, w存储边的权重
var idx int = 0
var st [N]bool  // 当前的点是否在队列当中。防止队列中存储重复的点
var dist [N]int // 存储所有点到1号点的距离
var cnt [N]int  // 当前最短路的边的个数

// queue
// 必须用无限长度的队列，如果用q[N],长度不够用，会报错
var q []int

func main() {
	// init h[]
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	for ; m > 0; m-- {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		add(a, b, c)
	}
	if spfa() {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}

func add(a, b, c int) {
	w[idx] = c
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func spfa() bool {
	for i := 1; i <= n; i++ {
		st[i] = true
		push(i)
	}
	for !empty() {
		t := front()
		pop()
		// 点从队列里出来了，不在队列里了，所以st[t] = false
		st[t] = false
		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]
				cnt[j] = cnt[t] + 1
				if cnt[j] >= n {
					return true
				}
				if !st[j] {
					push(j)
					st[j] = true
				}
			}
		}
	}
	return false
}

func push(x int) {
	q = append(q, x)
}

func pop() {
	q = q[1:]
}

func front() int {
	return q[0]
}

func empty() bool {
	if len(q) > 0 {
		return false
	}
	return true
}
