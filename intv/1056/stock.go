package main

import "fmt"

const N int = 1e5 + 10

var n int
var p [N]int // price数组，记录每天的股价
// g[i]表示在[1,i]中交易一次获得的最大收益
// f[i]表示在[i,n]交易一次获得的最大收益
var g, f [N]int

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &p[i])
	}
	// 求g[i]：在1 ~ i中买卖一次的最大收益
	// 这里i从2开始，是因为买卖一次至少需要2天
	for i, minv := 2, p[1]; i <= n; i++ {
		g[i] = max(g[i-1], p[i]-minv)
		minv = min(minv, p[i])
	}
	// 求f[i]：在i ~ n中买卖一次的最大收益
	// 这里i从n-1开始，也是因为买卖一次至少需要2天
	for i, maxv := n-1, p[n]; i >= 1; i-- {
		f[i] = max(f[i+1], maxv-p[i])
		maxv = max(maxv, p[i])
	}
	res := 0
	// 这里i从2开始，也是因为买卖一次至少需要2天
	for i := 2; i <= n; i++ {
		res = max(res, g[i]+f[i+1])
	}
	fmt.Printf("%d\n", res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
