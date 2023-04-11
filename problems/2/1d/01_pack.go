package main

import "fmt"

const N int = 1010

// 在二维版本中，f[i]只用到了f[i-1]这一层，所以可以用滚动数组来做, 直接把第一维删掉
var f [N]int    // 状态，一维优化版
var n, m int    // n表示物品个数，m表示背包容量
var v, w [N]int // v表示物品的体积，w表示物品的价值

// 一维优化版

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	fmt.Printf("%d\n", f[m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
