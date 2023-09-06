package main

import "fmt"

const N int = 310

var n, m int
var h [N][N]int // 棋盘高度
var f [N][N]int // 状态：所有从

// 4个方向
var dx = [4]int{1, -1, 0, 0}
var dy = [4]int{0, 0, 1, -1}

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scanf("%d", &h[i][j])
		}
	}
	// 递归算状态计算的套路
	// 先把每个状态初始化成-1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = -1
		}
	}
	res := 0
	// 枚举从每个点出发
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 这里不能直接写 f[i][j]，因为f[i][j]还没求出来
			res = max(res, dp(i, j))
		}
	}
	fmt.Printf("%d", res)
}

func dp(x, y int) int {
	// 如果f[i][j]已经算过了
	if f[x][y] != -1 {
		return f[x][y]
	}
	// 初始化f[x][y] = 1, 因为最次情况就是走当前这一个格子
	f[x][y] = 1
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a >= 1 && a <= n && b >= 1 && b <= m && h[a][b] < h[x][y] {
			f[x][y] = max(f[x][y], dp(a, b)+1)
		}
	}
	return f[x][y]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
