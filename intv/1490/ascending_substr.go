package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 10

var n int
var a, f, g [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	// 预处理 f[i]：以i结尾的单调上升子串的最大长度
	for i := 1; i <= n; i++ {
		if a[i-1] < a[i] {
			f[i] = f[i-1] + 1
		} else {
			f[i] = 1
		}
	}
	// 预处理 g[i]: 以i开头的单调上升子串的最大长度
	for i := n; i > 0; i-- {
		if a[i] < a[i+1] {
			g[i] = g[i+1] + 1
		} else {
			g[i] = 1
		}
	}
	res := 0
	// 枚举删除哪个数
	for i := 1; i <= n; i++ {
		if a[i-1] < a[i+1] {
			res = max(res, f[i-1]+g[i+1])
		} else {
			res = max(res, max(f[i-1], g[i+1]))
		}
	}
	fmt.Printf("%d\n", res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
