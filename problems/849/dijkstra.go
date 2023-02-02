package main

import "fmt"

const N int = 510

var n, m int
var g [N][N]int // graph,稠密图用邻接矩阵存储
var dist [N]int // 每个点到起点的最短距离
var st [N]bool  // 该点是否已经确定了最短距离

func main() {
	// init g[][] to infinity
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g[i][j] = 0x3f3f3f3f
		}
	}
	fmt.Scanf("%d%d", &n, &m)
	for ; m > 0; m-- {
		var x, y, z int
		fmt.Scanf("%d%d%d", &x, &y, &z)
		add(x, y, z)
	}
	d := dijkstra()
	fmt.Printf("%d", d)
}

func dijkstra() int {
	// init dist[] to infinity
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	dist[1] = 0              // 编号为1的点到起点的距离为0
	for i := 0; i < n; i++ { // 有n个点，每次循环找到一个点的最短路径，所以需要n次迭代
		//遍历 dist 数组，找到没有确定最短路径的节点中距离源点最近的点t
		t := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[j] < dist[t]) {
				t = j
			}
		}
		st[t] = true
		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], dist[t]+g[t][j])
		}
	}
	if dist[n] == 0x3f3f3f3f {
		return -1
	}
	return dist[n]
}

func add(x, y, z int) {
	//如果发生重边的情况则保留最短的一条边
	g[x][y] = min(g[x][y], z)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
