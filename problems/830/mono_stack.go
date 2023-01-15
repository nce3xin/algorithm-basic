package main

import "fmt"

const N int = 1e5 + 10

var stk []int = make([]int, N)
var tt int = 0 // tt是栈顶下标，栈从下标为1开始存

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)

		// tt > 0 stands for stk is not empty
		for tt > 0 && stk[tt] >= x {
			tt--
		}
		if tt > 0 {
			// stk is not empty
			fmt.Printf("%d ", stk[tt])
		} else {
			fmt.Printf("-1 ")
		}
		tt++
		stk[tt] = x
	}
}
