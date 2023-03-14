package main

import "fmt"

const (
	N   int = 210
	INF int = 1e9
)

var d [N][N]int
var n, m, q int

func main() {
	fmt.Scanf("%d%d%d", &n, &m, &q)
	// init d[][]
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = INF
			}
		}
	}
	for ; m > 0; m-- {
		var a, b, c int
		fmt.Scanf("%d%d%d", &a, &b, &c)
		// 重边取最短的那条
		d[a][b] = min(d[a][b], c)
	}
	floyd()
	for ; q > 0; q-- {
		var x, y int
		fmt.Scanf("%d%d", &x, &y)
		if d[x][y] > INF/2 {
			fmt.Printf("impossible\n")
		} else {
			fmt.Printf("%d\n", d[x][y])
		}
	}
}

func floyd() {
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
