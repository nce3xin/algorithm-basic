package main

import "fmt"

const N int = 1010
const MOD int = 1e9 + 7

var n int
var f [N][N]int

func main() {
	fmt.Scanf("%d", &n)
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			f[i][j] = (f[i-1][j-1] + f[i-j][j]) % MOD
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		res = (res + f[n][i]) % MOD
	}
	fmt.Printf("%d", res)
}
