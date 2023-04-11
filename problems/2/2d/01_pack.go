package main

import "fmt"

const N int = 1010

// 01 pack
var f [N][N]int // 状态, 集合：从前i个物品中选，总体积<=j的所有选法。集合的属性：所有选法中，价值的最大值
var n, m int    // n表示物品个数，m表示背包容量
var v, w [N]int // v表示物品的体积，w表示物品的价值

// 二维写法

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d", &v[i], &w[i])
	}
	// 初始化, 要枚举所有的状态，也就是f[0~n][0~m]
	// f[0][0~m] 表示的是考虑0件物品，总体积不超过0，不超过1…一直到不超过m，这样的情况下，最大价值是多少。
	// 由于1件物品都没选，所以f[0][0~m]都是0. 由于f数组定义成了全局变量，它默认就是0。
	// 所以f[0][0~m]=0 这句初始化就可以不写
	// 所以i从1开始枚举（因为f[0][0~m]都已经全部是0了）
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i-1][j-v[i]]+w[i])
			}
		}
	}
	fmt.Printf("%d\n", f[n][m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
