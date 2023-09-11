package main

import "fmt"

const N int = 1010
const M int = N * N // 队列长度上限

type PII struct {
	x int
	y int
}

var n, m int
var g [N]string   // 矩形土地
var q [M]PII      // 队列
var st [N][N]bool // 记录某个点是否访问过

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &g[i])
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] == 'W' && !st[i][j] {
				bfs(i, j)
				cnt++
			}
		}
	}
	fmt.Printf("%d", cnt)
}

// sx: start x
// sy: start y
func bfs(sx, sy int) {
	// 队列的队头和队尾
	hh := 0
	tt := 0
	// 一开始队列里只有一个元素
	q[tt] = PII{
		x: sx,
		y: sy,
	}
	// 防止重复遍历某个点
	st[sx][sy] = true
	// 八连通，习惯写双层循环，遍历3*3的格子，再把中间那个格子挖掉
	for hh <= tt {
		t := q[hh]
		hh++
		for i := t.x - 1; i <= t.x+1; i++ {
			for j := t.y - 1; j <= t.y+1; j++ {
				// 把中间的格子挖掉
				if i == t.x && j == t.y {
					continue
				}
				if i < 0 || i >= n || j < 0 || j >= m {
					continue
				}
				if g[i][j] == '.' || st[i][j] {
					continue
				}
				tt++
				q[tt] = PII{
					x: i,
					y: j,
				}
				st[i][j] = true
			}
		}
	}
}
