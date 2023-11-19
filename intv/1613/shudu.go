package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 10

// 这儿不能写成 var g [N][N]byte, 如果规定了第二维的长度，在读入input的时候会有问题
var g = make([][]byte, N)

// row表示第几行有没有填过某个数，例如row[1][2]表示第1行有没有填过数字2
// col意义和row类似
var row, col [N][N]bool

// 3*3格子内有没有填过某个数
var cell [N][N][N]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < 9; i++ {
		fmt.Fscanln(in, &g[i])
		for j := 0; j < 9; j++ {
			if g[i][j] != '.' {
				t := g[i][j] - '0'
				row[i][t] = true
				col[j][t] = true
				cell[i/3][j/3][t] = true
			}
		}
	}
	dfs(0, 0)
}

// dfs函数有返回值是因为题目要求只要求一种填法。所以找到一个解法的时候，就可以直接return
func dfs(x, y int) bool {
	if y == 9 {
		return dfs(x+1, 0)
	}
	if x == 9 {
		for i := 0; i < 9; i++ {
			fmt.Printf("%s\n", g[i])
		}
		return true
	}

	// 这句别忘了，否则打印出来的答案是空的
	if g[x][y] != '.' {
		return dfs(x, y+1)
	}

	// 枚举当前位置能填哪些数
	for i := 1; i <= 9; i++ {
		if !row[x][i] && !col[y][i] && !cell[x/3][y/3][i] {
			g[x][y] = byte(i) + '0'
			row[x][i] = true
			col[y][i] = true
			cell[x/3][y/3][i] = true
			if dfs(x, y+1) {
				return true
			}
			g[x][y] = '.'
			row[x][i] = false
			col[y][i] = false
			cell[x/3][y/3][i] = false
		}
	}
	return false
}
