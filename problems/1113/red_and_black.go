package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 25

var n, m int // 行数，列数
var g [N]string
var st [N][N]bool

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Fscan(in, &m, &n)
		// input是两个零，表示输入结束
		if m == 0 && n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &g[i])
		}
		// 找一下起点在什么地方
		var x, y int
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if g[i][j] == '@' {
					x = i
					y = j
				}
			}
		}

		// re-init st
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				st[i][j] = false
			}
		}

		// dfs()函数返回的就是它能搜到的新大陆的数量
		fmt.Printf("%d\n", dfs(x, y))
	}
}

// dfs()函数返回的就是它能搜到的新大陆的数量
func dfs(x, y int) int {
	// 最开始的时候有1个点，就是当前这个点
	cnt := 1
	st[x][y] = true
	for i := 0; i < 4; i++ {
		a := x + dx[i]
		b := y + dy[i]
		if a < 0 || a >= n || b < 0 || b >= m {
			continue
		}
		if g[a][b] != '.' { // 不是黑砖
			continue
		}
		if st[a][b] {
			continue
		}
		cnt += dfs(a, b)
	}
	return cnt
}
