package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 10

var n int
var nums [N]int
var st [N]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	dfs(0, a)
}

func dfs(u int, a []int) {
	if u == n {
		// 输出必须要换成这种写法，用fmt.Printf()会TLE
		out := bufio.NewWriter(os.Stdout)
		defer out.Flush() // 必须要写这句，不然没输出
		for i := 0; i < n; i++ {
			fmt.Fprintf(out, "%d ", nums[i])
		}
		fmt.Fprintln(out)
		return
	}
	for i := 0; i < n; i++ {
		if !st[i] {
			st[i] = true
			nums[u] = a[i]
			dfs(u+1, a)
			st[i] = false
			// 去重，一旦说当前的数已经用过了，那么就需要把它后面和它相同的数全部跳过
			for i+1 < n && a[i] == a[i+1] {
				i++
			}
		}
	}
}
