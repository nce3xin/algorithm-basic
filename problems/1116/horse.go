package main

import "fmt"

const N int = 10

var st [N][N]bool
var n, m int
var ans int

var dx = [8]int{-2, -1, 1, 2, 2, 1, -1, -2}
var dy = [8]int{1, 2, 2, 1, -1, -2, -2, -1}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var x, y int
		fmt.Scanf("%d%d%d%d", &n, &m, &x, &y)
		ans = 0
		dfs(x, y, 1)
		fmt.Printf("%d\n", ans)
	}
}

// cnt: 当前正在搜第几个点
func dfs(x, y, cnt int) {
	if cnt == n*m { // 如果正在搜的点是最后一个点
		ans++
		return
	}
	st[x][y] = true

	// 枚举8个方向
	for i := 0; i < 8; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a < 0 || a >= n || b < 0 || b >= m {
			continue
		}
		if st[a][b] {
			continue
		}

		dfs(a, b, cnt+1)
	}
	// 恢复现场
	st[x][y] = false
}
