package main

import "fmt"

const N int = 100 + 10

var g [N][N]int  // graph
var d [N][N]int  // 每个点到起点的距离
var q [N * N]PII // manual queue for BFS
var n, m int
var prev [N][N]PII // 记一下每个点的前一个点是哪个

type PII struct {
	first  int
	second int
}

func main() {
	fmt.Scanf("%d%d", &n, &m)
	// read graph
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &g[i][j])
		}
	}
	// init d[][]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			d[i][j] = -1
		}
	}

	fmt.Printf("%d", bfs())
}

func bfs() int {
	hh, tt := 0, -1 // for queue
	// 起点入队
	tt++
	q[tt] = PII{0, 0}
	d[0][0] = 0 // 表示起点走过了

	// 4 directions
	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	for hh <= tt { // queue is not empty
		// 取出队头元素
		t := q[hh]
		hh++
		// 扩展元素
		for i := 0; i < 4; i++ { // 4个方向
			x := t.first + dx[i]
			y := t.second + dy[i]
			// d[x][y] == -1 表示点 (x,y) 还没有走过
			if x >= 0 && x < n && y >= 0 && y < m && g[x][y] == 0 && d[x][y] == -1 {
				tt++
				q[tt] = PII{x, y}
				d[x][y] = d[t.first][t.second] + 1
				// 记录当前扩展的点是从哪个点来的
				prev[x][y] = t
			}
		}
	}

	// 打印最短路径
	x := n - 1
	y := m - 1
	for x != 0 || y != 0 { // 没到起点，就一直循环
		fmt.Printf("%d %d\n", x, y)
		t := prev[x][y]
		x = t.first
		y = t.second
	}
	return d[n-1][m-1]
}
