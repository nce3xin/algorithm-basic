package main

import "fmt"

const N int = 1010

var n, m int    // n表示物品个数，m表示背包容量
var f [N]int    // 状态, 集合
var v, w [N]int // v表示物品的体积，w表示物品的价值

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d%d", &v[i], &w[i])
	}
	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			f[j] = f[j]
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	fmt.Printf("%d", f[m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
