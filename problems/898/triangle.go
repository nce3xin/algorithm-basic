package main

import "fmt"

const N int = 510
const INF int = 1e9

var n int       // 三角形的层数
var a [N][N]int // 三角形中的数值
var f [N][N]int // f[i][j]表示的集合: 所有从起点，走到(i,j)的路径

// 这里是从上到下计算，也可以从下到上，终点就是第一行的点，只有一个，不用遍历了，代码会短一些
func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ { // i从1开始是因为涉及到i-1操作
		for j := 1; j <= i; j++ { // j从1开始是因为涉及到j-1操作
			fmt.Scanf("%d", &a[i][j])
		}
	}
	// 为了不处理边界，先把所有f都置为负无穷
	for i := 0; i <= n; i++ {
		// j<=i+1,要加1，是因为算最右边点的时候，要用到它的右上角的点，但这个点不存在，所有需要初始化成负无穷
		// 也就是每行多初始化一个
		for j := 0; j <= i+1; j++ {
			f[i][j] = -INF
		}
	}
	f[1][1] = a[1][1]
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			f[i][j] = max(f[i-1][j-1]+a[i][j], f[i-1][j]+a[i][j])
		}
	}
	// 最后答案要遍历最后一行，因为最后一行有可能走到每一个位置
	// 把每一个位置的f[i][j]取一个最大值
	res := -INF
	for i := 1; i <= n; i++ {
		res = max(res, f[n][i])
	}
	fmt.Printf("%d", res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
