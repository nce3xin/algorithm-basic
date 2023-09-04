package main

import "fmt"

const N int = 1010

var n int
var a, f [N]int

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ { // 因为涉及到i-1操作，所有i从1开始
		fmt.Scanf("%d", &a[i])
	}
	for i := 1; i <= n; i++ {
		// 最极端的情况，以i结尾的最长上升子序列长度是1。只有i一个数，所以f[i]起码是等于1的
		// 只有a[i]一个数
		f[i] = 1
		for j := 1; j <= n; j++ {
			if a[j] < a[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		res = max(res, f[i])
	}
	fmt.Printf("%d", res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
