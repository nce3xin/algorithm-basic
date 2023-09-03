package main

import "fmt"

const N int = 110

var n, m int       // 物品组数和背包容量
var v, w [N][N]int // 物品的体积、价值
var f [N]int       // 直接写成一维
var s [N]int       // 每组物品个数

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &s[i])
		for j := 0; j < s[i]; j++ {
			fmt.Scanf("%d%d", &v[i][j], &w[i][j])
		}
	}
	for i := 1; i <= n; i++ {
		for j := m; j >= 0; j-- {
			for k := 0; k < s[i]; k++ {
				if v[i][k] <= j {
					f[j] = max(f[j], f[j-v[i][k]]+w[i][k])
				}
			}
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
