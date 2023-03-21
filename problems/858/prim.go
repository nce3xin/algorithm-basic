package main

import "fmt"

const N int = 510

// 稠密图
var g [N][N]int
var n, m int
var dist [N]int          // 表示当前点到集合的距离，而不是到1号点的距离
var st [N]bool           // 表示已经找到最短距离的点的集合
var INF int = 0x3f3f3f3f // 最大值

func main() {
	fmt.Scanf("%d%d", &n, &m)
	// init g
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g[i][j] = INF
		}
	}
	for ; m > 0; m-- {
		var a, b, c int
		fmt.Scanf("%d%d%d", &a, &b, &c)
		g[a][b] = min(g[a][b], c)
		g[b][a] = g[a][b]
	}

	t := prim()
	if t == INF {
		fmt.Printf("impossible\n")
	} else {
		fmt.Printf("%d\n", t)
	}
}

func prim() int {
	res := 0
	// init dist
	for i := 0; i < N; i++ {
		dist[i] = INF
	}
	for i := 0; i < n; i++ {
		t := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		if i != 0 && dist[t] == INF {
			return INF
		}
		// 需要先累加，再更新,否则自环会出现问题
		if i != 0 {
			res += dist[t]
		}
		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], g[t][j])
		}
		st[t] = true
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
