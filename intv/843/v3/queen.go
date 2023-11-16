package main

import "fmt"

const N int = 15

var n int
var g [N][N]byte
var col [N]bool
var dg, udg [N * 2]bool

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = '.'
		}
	}
	dfs(0)
}

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%s", string(g[i][j]))
			}
			fmt.Println()
		}
		fmt.Println()
		return
	}
	// 枚举每一列
	for i := 0; i < n; i++ {
		// 和v1的区别，只有udg[u-i+n]不一样，下标不管是u-i+n还是 i-u+n都可以
		if !col[i] && !dg[u+i] && !udg[u-i+n] {
			col[i] = true
			dg[u+i] = true
			udg[u-i+n] = true
			g[u][i] = 'Q'
			dfs(u + 1)
			g[u][i] = '.'
			col[i] = false
			dg[u+i] = false
			udg[u-i+n] = false
		}
	}
}
