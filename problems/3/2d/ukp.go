package main

import "fmt"

const N int = 1010

var n, m int    // n表示物品个数，m表示背包容量
var f [N][N]int // 状态, 集合：从前i个物品中选，总体积<=j的所有选法。集合的属性：所有选法中，价值的最大值
var v, w [N]int // v表示物品的体积，w表示物品的价值

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i][j-v[i]]+w[i])
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
