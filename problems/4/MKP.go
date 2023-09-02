package main

import "fmt"

const N int = 110

var n, m int
var v, w, s [N]int
var f [N][N]int // 状态, 集合：从前i个物品中选，总体积<=j的所有选法。集合的属性：所有选法中，价值的最大值

// 多重背包问题的暴力解法
func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d%d", &v[i], &w[i], &s[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			for k := 0; k <= s[i] && k*v[i] <= j; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}
	fmt.Printf("%d", f[n][m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
