package main

import "fmt"

const N int = 110

var n, m int
var res [N][N]int

// 按照 右->下->左->上 顺序来定
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func main() {
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	k := 1 // 当前要放置的数
	x := 0 // 坐标
	y := 0 // 坐标
	d := 0 // 方向
	for ; k <= n*m; k++ {
		res[x][y] = k // 放置数k
		// 计算下一个坐标
		a := x + dx[d]
		b := y + dy[d]
		// 判断下一个坐标有没有出界，或者被访问过
		if a < 0 || a >= n || b < 0 || b >= m || res[a][b] != 0 {
			// 用取余数来循环方向
			d = (d + 1) % 4
			// 修改下一个坐标
			a = x + dx[d]
			b = y + dy[d]
		}
		// 把下一个坐标赋值给(x,y)坐标
		x = a
		y = b
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Printf("\n")
	}
}
