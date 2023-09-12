package main

import "fmt"

const N int = 110

var n int
var g [N]string
var st [N][N]bool
var xa, ya, xb, yb int

var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		fmt.Scanf("%d", &n)
		for i := 0; i < n; i++ {
			// g[i]前面一定要加&, 要不然读不进来数！会一直报index out of range!
			fmt.Scanf("%s", &g[i])
		}
		fmt.Scanf("%d%d%d%d", &xa, &ya, &xb, &yb)

		// re-init st, 也是必须的
		// 因为st是全局数组，在上一个testcase里已经被修改过了, 需要重新初始化
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				st[i][j] = false
			}
		}
		if dfs(xa, ya) {
			fmt.Printf("YES\n")
		} else {
			fmt.Printf("NO\n")
		}
	}
}

// 表示能不能从x,y点走到xb, yb点
func dfs(x, y int) bool {
	if g[x][y] == '#' {
		return false
	}
	if x == xb && y == yb {
		return true
	}
	st[x][y] = true
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a < 0 || a >= n || b < 0 || b >= n {
			continue
		}
		if st[a][b] {
			continue
		}

		if dfs(a, b) {
			return true
		}
	}
	return false
}
