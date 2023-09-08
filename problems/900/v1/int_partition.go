package main

import "fmt"

// 转化成完全背包问题

const N int = 1010
const MOD int = 1e9 + 7

var n int
var f [N]int

func main() {
	fmt.Scanf("%d", &n)
	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			f[j] = (f[j] + f[j-i]) % MOD
		}
	}
	fmt.Printf("%d", f[n])
}
