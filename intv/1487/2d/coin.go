package main

import "fmt"

const N int = 110
const M int = 1e5 + 10
const MOD int = 1e9 + 7

var n1, n2, m int

// 所有只从前i个硬币中选，且总面值是j的选法的数量
var f [N][M]int

// 二维写法

func main() {
	fmt.Scanf("%d%d%d", &n1, &n2, &m)
	// f[0][0]: 从前0个硬币中选，总面值为0，有一种选法，所以f[0][0] = 1
	// f[0][...]: 其他情况，从前0个硬币中选，总面值为非0，没有这样的选法，所以f[0][...] = 0
	f[0][0] = 1
	// 普通硬币，每种可以取无限个，完全背包问题
	for i := 1; i <= n1; i++ {
		var p int
		fmt.Scanf("%d", &p)
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= p {
				f[i][j] = (f[i][j] + f[i][j-p]) % MOD
			}
		}
	}
	// 纪念币，每种最多只能取一个，01背包问题
	// 这儿不能写成 i := 1; i <= n1; i++，因为这样f[i][j]从n1之后就一直没计算了
	for i := n1 + 1; i <= n1+n2; i++ {
		var p int
		fmt.Scanf("%d", &p)
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= p {
				f[i][j] = (f[i][j] + f[i-1][j-p]) % MOD
			}
		}
	}
	fmt.Printf("%d\n", f[n1+n2][m])
}
