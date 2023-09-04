package main

import "fmt"

const N int = 1010

var n, m int
var a, b string
var f [N][N]int // 状态表示，f[i][j]表示a的前i个字母，和b的前j个字母的最长公共子序列长度

func main() {
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i-1][j], f[i][j-1])
			if a[i-1] == b[j-1] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}
	fmt.Printf("%d", f[n][m])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
