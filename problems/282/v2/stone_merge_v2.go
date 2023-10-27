package main

import "fmt"

const N int = 310
const INF int = 1e9

var n int
var a, s [N]int // a存数组，s存前缀和
var f [N][N]int // f[i][j]表示合并区间[i,j]的最小代价

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ { // 从下标为1开始读入
		fmt.Scanf("%d", &a[i])
	}
	// init f[][]
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			f[i][j] = INF
		}
	}
	// 算前缀和
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i]
		// i自己一堆，不用合并，代价为0
		f[i][i] = 0
	}
	// 从区间长度开始枚举，区间[i,j]的长度是j-i
	// 如果结果哪里不对，比如等于1e9 INF, 就是3个for循环的条件里，小于号或者小于等于号的差别，都改改试试
	for length := 1; length < n; length++ {
		// 枚举区间左端点
		for i := 1; i <= n-length; i++ { // 不能是i <= n-length +1, 否则i+length就是n+1,会爆掉
			// 枚举分界点k: 左闭右开
			for k := i; k < i+length; k++ {
				f[i][i+length] = min(f[i][i+length], f[i][k]+f[k+1][i+length]+s[i+length]-s[i-1])
			}
		}
	}
	fmt.Printf("%d\n", f[1][n])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
