package main

import "fmt"

const N int = 110
const M int = 20

// f[i,j] 表示所有用j个鸡蛋测i次的测量方案的集合。属性：最多能测量多长的区间。
// i表示测几次，n层楼的最多测n次，所以i的上界也是N
var f [N][M]int

var n, m int

func main() {
	for {
		_, err := fmt.Scanf("%d%d", &n, &m)
		if err != nil {
			break
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				f[i][j] = f[i-1][j-1] + f[i-1][j] + 1
			}
			if f[i][m] >= n {
				fmt.Printf("%d\n", i)
				break
			}
		}
	}
}
