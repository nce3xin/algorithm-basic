package main

import "fmt"

const N int = 55
const M int = N * N

var n, m int
var g [N][N]int // 图
var q [M]PII    // 队列
var st [N][N]bool

type PII struct {
	x int
	y int
}

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &g[i][j])
		}
	}
	// 需要统计两个值
	cnt := 0
	area := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !st[i][j] {
				// 假设bfs有个返回值，是连通块的面积
				area = max(area, bfs(i, j))
				cnt++ // 连通块个数+1
			}
		}
	}

	fmt.Printf("%d\n", cnt)
	fmt.Printf("%d\n", area)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// sx: start x
// sy: start y
func bfs(sx, sy int) int {
	// 为了方便取二进制位，这里的dx,dy按照题中1,2,4,8的顺序来，也就是左，上，右，下
	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}
	hh := 0
	tt := 0
	area := 0
	q[tt] = PII{
		x: sx,
		y: sy,
	}
	st[sx][sy] = true
	// 当队列不空时
	for hh <= tt {
		t := q[hh]
		hh++
		// 出队时统计area，其实入队时统计也可以
		area++

		for i := 0; i < 4; i++ {
			a := t.x + dx[i]
			b := t.y + dy[i]
			if a < 0 || a >= n || b < 0 || b >= m {
				continue
			}
			if st[a][b] {
				continue
			}
			// 判断连通性
			// 判断t点在i这个方向上是不是1，是的话是墙
			if g[t.x][t.y]>>i&1 == 1 {
				continue
			}
			tt++
			q[tt] = PII{
				x: a,
				y: b,
			}
			st[a][b] = true
		}
	}
	return area
}
