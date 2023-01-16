package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int64 = 1e6 + 10

var a []int = make([]int, N)
var q []int = make([]int, N) // 用队列维护窗口，队列存的不是值，是下标
var hh, tt int = 0, -1

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(in, &n, &k)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 0; i < n; i++ {
		// 判断队头元素是否已超出窗口范围
		// hh <= tt 表示队列不为空
		// 这里用if不用for，是因为每次只会向右移动一位
		if hh <= tt && q[hh] < i-k+1 {
			hh++
		}
		// hh <= tt 表示队列不为空
		// 求窗口最小值，如果存在逆序对，一直删，直到没有逆序对
		for hh <= tt && a[q[tt]] >= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		// 只有窗口元素数量够k个数时，才输出
		if i >= k-1 {
			fmt.Printf("%d ", a[q[hh]])
		}
	}

	fmt.Printf("\n")

	hh, tt = 0, -1 // 记得重新初始化
	// 求窗口最大值，和上面完全对称的写法
	for i := 0; i < n; i++ {
		// 判断队头元素是否已超出窗口范围
		// hh <= tt 表示队列不为空
		if hh <= tt && q[hh] < i-k+1 {
			hh++
		}
		for hh <= tt && a[q[tt]] <= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		// 只有窗口元素数量够k个数时，才输出
		if i >= k-1 {
			fmt.Printf("%d ", a[q[hh]])
		}
	}
}
