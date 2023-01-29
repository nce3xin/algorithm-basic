package main

import "fmt"

const N int = 20 //对角线条数 2n-1 取20防止越界

var n int
var g [N][N]byte // board
// bool数组用来判断搜索的下一个位置是否可行
// col列，dg对角线，udg反对角线是否存在皇后，下标表示列、对角线、反对角线的编号
// 例如，dg[3]就表示编号为3的对角线上是否存在皇后
var col, dg, udg [N]bool

func main() {
	fmt.Scanf("%d", &n)

	// init board
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g[i][j] = '.'
		}
	}

	// 枚举第0行
	dfs(0)
}

// 枚举第u行
func dfs(u int) {
	//放满了棋盘，输出棋盘
	if u == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%c", g[i][j])
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
		return
	}
	// 枚举第u行的所有列
	for i := 0; i < n; i++ {
		if !col[i] && !dg[u+i] && !udg[i-u+n] { //不冲突，放皇后
			g[u][i] = 'Q'
			//对应的列、斜线，状态改变
			col[i] = true
			dg[u+i] = true
			udg[i-u+n] = true
			dfs(u + 1) //处理下一行
			g[u][i] = '.'
			//恢复现场
			col[i] = false
			dg[u+i] = false
			udg[i-u+n] = false
		}
	}
}
