package main

import "fmt"

const M int = 1e5 + 10
const MOD int = 1e9 + 7

var n1, n2, m int

// 所有只从前i个硬币中选，且总面值是j的选法的数量
var f [M]int

// 一维写法

func main() {
	fmt.Scanf("%d%d%d", &n1, &n2, &m)
	// f[0][0]: 从前0个硬币中选，总面值为0，有一种选法，所以f[0][0] = 1
	// f[0][...]: 其他情况，从前0个硬币中选，总面值为非0，没有这样的选法，所以f[0][...] = 0
	f[0] = 1
	for i := 1; i <= n1; i++ {
		var p int
		fmt.Scanf("%d", &p)
		for j := p; j <= m; j++ {
			f[j] = (f[j] + f[j-p]) % MOD
		}
	}
	// 这儿不能写成 i := 1; i <= n1; i++，因为这样f[i][j]从n1之后就一直没计算了
	for i := n1 + 1; i <= n1+n2; i++ {
		var p int
		fmt.Scanf("%d", &p)
		for j := m; j >= p; j-- {
			f[j] = (f[j] + f[j-p]) % MOD
		}
	}
	fmt.Printf("%d\n", f[m])
}
