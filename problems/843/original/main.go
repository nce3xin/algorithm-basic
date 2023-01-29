package main

import "fmt"

const N int = 20

var n int
var g [N][N]byte // board
var row, col, dg, udg [N]bool

func main() {
	fmt.Scanf("%d", &n)
	// init board
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = '.'
		}
	}

	dfs(0, 0, 0)
}

// dfs x: 行，y: 列，s: 已经放置的皇后个数
// dfs 枚举当前位置要不要放皇后
func dfs(x, y, s int) {
	if y == n { // 列越界了，给它转回来到下一行的开头
		y = 0
		x++
	}
	if x == n { // 行越界了，代表所有行都已经遍历结束
		if s == n { // 已经放置的皇后个数等于要求放置的皇后个数，找到一个解，输出
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					fmt.Printf("%c", g[i][j])
				}
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
		return
	}

	// 当前位置不放皇后
	dfs(x, y+1, s)

	// 当前位置放皇后
	if !row[x] && !col[y] && !dg[x+y] && !udg[x-y+n] { // 当前位置的行、列、对角线都没有皇后
		row[x] = true
		col[y] = true
		dg[x+y] = true
		udg[x-y+n] = true
		g[x][y] = 'Q'
		dfs(x, y+1, s+1)
		// 恢复现场
		g[x][y] = '.'
		row[x] = false
		col[y] = false
		dg[x+y] = false
		udg[x-y+n] = false
	}
}
