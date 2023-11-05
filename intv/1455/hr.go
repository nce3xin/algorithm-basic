package main

import "fmt"

const M int = 1010

var T int
var a [M]int // 存黑板上的若干个数字

// f[n,k]: 表示0~n-1这n个人围成一圈，从第k个数A[k]开始用，最终剩余的编号是多少

func main() {
	fmt.Scanf("%d", &T)
	for ; T > 0; T-- {
		var n, m int
		fmt.Scanf("%d%d", &n, &m)
		for i := 0; i < m; i++ {
			fmt.Scanf("%d", &a[i])
		}
		// 因为算f[n,1]只会用到一个数，就是 f[n-1,2], 所以可以用一个变量来存
		// 就不用开f[]数组了
		res := 0
		// 每次是用i去推i+1,所以i只用循环到n-1. 因为用n-1就能推出n了
		for i, j := 1, (n-1)%m; i < n; {
			// f[n,k] = (f[n-1,k+1] + A_k) % n，从右边反推前面，所以需要i++,j减1再对m取模
			i++
			// 这里本来是j = (j-1) % m, 之所以要加上m，是因为j-1可能为负数。在对m取模的意义下加一个m并不影响结果
			j = (j + m - 1) % m
			res = (res + a[j]) % i
		}
		fmt.Printf("%d\n", res)
	}
}
