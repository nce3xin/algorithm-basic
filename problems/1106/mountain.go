package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1010
const M int = N * N

var n int
var h [N][N]int
var q [M]PII
var st [N][N]bool

type PII struct {
	x int
	y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &h[i][j])
		}
	}
	peak := 0
	valley := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !st[i][j] {
				hasHigher := false
				hasLower := false
				bfs(i, j, &hasHigher, &hasLower)
				if !hasHigher {
					peak++
				}
				if !hasLower {
					valley++
				}
			}
		}
	}
	fmt.Printf("%d %d", peak, valley)
}

func bfs(sx, sy int, hasHigher *bool, hasLower *bool) {
	hh := 0
	tt := 0
	q[tt] = PII{
		x: sx,
		y: sy,
	}
	st[sx][sy] = true
	for hh <= tt {
		t := q[hh]
		hh++
		// 八连通，用双层循环，再把中间的格子挖掉
		for i := t.x - 1; i <= t.x+1; i++ {
			for j := t.y - 1; j <= t.y+1; j++ {
				// 扣掉中间的格子，这句不写也对，因为会在 else !st[i][j] 判断时会被判掉
				if i == t.x && j == t.y {
					continue
				}
				if i < 0 || i >= n || j < 0 || j >= n {
					continue
				}
				if h[i][j] != h[t.x][t.y] { // 山脉的边界
					if h[i][j] > h[t.x][t.y] {
						*hasHigher = true
					} else { // 这里else只会是小于的情况，不会等于，因为外层还有一个if判断
						*hasLower = true
					}
				} else if !st[i][j] {
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
}
