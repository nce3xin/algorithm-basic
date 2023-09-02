package main

import "fmt"

const N int = 25000 // 拆分之后的物品总个数：N*logS = 1000 * log(2000) = 1000*12 = 24000，再打点提前量
const M int = 2010  // 背包体积的上限

var n, m int    // 用来记录物品个数N和背包体积V
var v, w [N]int // 第i种物品的体积、价值, 不用定义数量
var f [N]int    // 01背包可以优化成1维，这里直接写成1维了

func main() {
	fmt.Scanf("%d%d", &n, &m)
	cnt := 0 // 存储所有新的物品，表示所有新的物品的编号
	for i := 1; i <= n; i++ {
		var a, b, s int //第i个物品的体积，价值，数量
		fmt.Scanf("%d%d%d", &a, &b, &s)
		k := 1
		for k <= s {
			cnt++          // 物品编号+1
			v[cnt] = a * k // a是第i个物品体积，第i个物品打包k个，新的物品体积就是a*k
			w[cnt] = b * k // b是第i个物品价值，第i个物品打包k个，新的物品体积就是b*k
			s -= k         // 这句和下面那句顺序不能反
			k = k * 2
		}
		if s > 0 { // s > 0 说明物品还剩下一些需要补上
			cnt++
			v[cnt] = a * s
			w[cnt] = b * s
		}
	}

	n = cnt // n更新成拆分后的物品总数

	// 然后对拆分后的所有物品做一次01背包问题，这里直接写成1维的
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
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
