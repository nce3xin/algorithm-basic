package main

// 用滚动数组优化空间复杂度

import "fmt"

const (
	MOD int = 1e9 + 7
	N   int = 5010
	M   int = 8192
)

var n int
var a [N]int

// 从前i个数中选且异或和为j的所有方案的集合
var f [2][M]int

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	// 当有0个数的时候，异或和是0，有1个方案. f[0][1], f[0][2]... 都是0
	f[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j < M; j++ {
			f[i&1][j] = f[(i-1)&1][j]
			if j^a[i] < M {
				f[i&1][j] = (f[i&1][j] + f[(i-1)&1][j^(a[i])]) % MOD
			}
		}
	}
	res := 0
	for i := 2; i < M; i++ {
		if isPrime(i) {
			res = (res + f[n&1][i]) % MOD
		}
	}
	fmt.Printf("%d\n", res)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
