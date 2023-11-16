package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1010

var n, m int
var f [N][N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var w int
			fmt.Fscan(in, &w)
			// w就是当前[i,j]位置上的值，只有它是1，f[i][j]才有可能有全部包含1的正方形
			if w != 0 {
				f[i][j] = min(f[i-1][j], min(f[i][j-1], f[i-1][j-1])) + 1
				res = max(res, f[i][j])
			}
		}
	}
	// 注意求的是面积，我们记录的是边长，所以最后要算平方
	fmt.Printf("%d\n", res*res)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
