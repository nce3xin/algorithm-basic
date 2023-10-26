package main

import "fmt"

const N int = 110
const M int = 20

var f [N][M]int // i层楼，j个鸡蛋的测量方案中最坏情况的最小值

func main() {
	var n, m int
	for {
		_, err := fmt.Scanf("%d%d", &n, &m)
		// 输入结束
		if err != nil {
			break
		}
		// init
		for i := 1; i <= n; i++ {
			// 用1个鸡蛋去测长度为i的楼，肯定是需要测i次的
			f[i][1] = i
		}
		for j := 1; j <= m; j++ {
			// 如果高度是1，有i个鸡蛋，那也是只需要测1次
			f[1][j] = 1
		}

		// i,j取1已经初始化过了，所以这里i,j都从2开始遍历
		for i := 2; i <= n; i++ {
			for j := 2; j <= m; j++ {
				f[i][j] = f[i][j-1]
				for k := 1; k <= i; k++ {
					// 最后的加1，是要加上第k层摔鸡蛋的这一次
					f[i][j] = min(f[i][j], max(f[k-1][j-1], f[i-k][j])+1)
				}
			}
		}
		fmt.Printf("%d\n", f[n][m])
	}
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
