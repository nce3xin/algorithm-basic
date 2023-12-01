package main

const INF = 1e8
const N int = 1e4 + 10

// 这题不适合写2d，初始化条件不好写，而且2d版本提交会TLE

func coinChange(coins []int, amount int) int {
	// 完全背包问题
	// f[i][j]: 从前i种硬币中选，能凑出总金额j的所有选法。属性：所有选法中硬币数量的最小值
	var f [N]int
	for i := 0; i < N; i++ {
		f[i] = INF
	}
	n := len(coins)
	f[0] = 0
	for i := 1; i <= n; i++ {
		for j := coins[i-1]; j <= amount; j++ {
			f[j] = min(f[j], f[j-coins[i-1]]+1)
		}
	}
	if f[amount] == INF {
		return -1
	}
	return f[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
