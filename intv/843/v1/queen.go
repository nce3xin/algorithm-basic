package main

import "fmt"

const N int = 15

var n int
var g [N][N]byte
var col [N]bool
var dg, udg [N * 2]bool // 对角线数量大概是n的两倍

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = '.'
		}
	}
	// 按行枚举
	dfs(0)
}

func dfs(u int) {
	if u == n { // 所有行枚举完了，找到了一个答案
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
		// i-u可能是负数，所以加个n变成正的
		if !col[i] && !dg[u+i] && !udg[i-u+n] {
			col[i] = true
			dg[u+i] = true
			udg[i-u+n] = true
			g[u][i] = 'Q'
			dfs(u + 1)
			g[u][i] = '.'
			col[i] = false
			dg[u+i] = false
			udg[i-u+n] = false
		}
	}
}
