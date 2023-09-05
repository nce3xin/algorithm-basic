package main

import "fmt"

const N int = 310
const INF int = 1e9

var n int
var s [N]int    // 前缀和
var f [N][N]int // 集合：所有将第i堆石子到第j堆石子合并成一堆石子的合并方式。属性：代价最小值

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	// 求前缀和，这里公式和之前讲的不一样，是因为把数组直接存进了s[i]
	for i := 1; i <= n; i++ {
		s[i] += s[i-1]
	}
	// 按照区间长度从小到大来枚举所有状态
	for length := 1; length <= n; length++ {
		// 接着枚举左端点（右端点由左端点和区间长度去确定）
		for i := 2; i+length-1 <= n; i++ {
			l := i
			r := i + length - 1 //计算右端点
			f[l][r] = INF       //初始化f[i][j]
			// 最后枚举分段点k
			for k := l; k < r; k++ {
				f[l][r] = min(f[l][r], f[l][k]+f[k+1][r]+s[r]-s[l-1])
			}
		}
	}
	fmt.Printf("%d", f[1][n])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
