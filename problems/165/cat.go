package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const N int = 20

var n int       // 猫的数量
var m int       // 每量缆车的最大承重量
var w [N]int    // 每个小猫的重量
var sum [N]int  // 每辆车上当前的总重量
var ans int = N // 结果

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &w[i])
	}
	// 优化搜索顺序，猫按重量从大到小排序，先搜重量大的猫
	wSlice := w[:]                                 // 数组转换成切片
	sort.Sort(sort.Reverse(sort.IntSlice(wSlice))) // 降序排列
	// 当前搜到第0只猫，当前缆车的数量是0
	dfs(0, 0)
	fmt.Printf("%d\n", ans)
}

// u: 当前第几只小猫, 是下标。闫总的习惯，u表示下标
// k: 当前缆车数量
func dfs(u, k int) {
	// 最优性剪枝
	if k >= ans {
		return
	}
	// 如果已经搜完n只小猫
	if u == n {
		// 这里k一定严格小于ans，因为如果k>=ans, 上面的代码就会直接return
		ans = k
		return
	}
	// 枚举当前小猫可以放在哪辆车上
	for i := 0; i < k; i++ {
		// 可行性剪枝
		if sum[i]+w[u] <= m {
			// 把小猫放到第i辆车上
			sum[i] += w[u]
			dfs(u+1, k)
			sum[i] -= w[u] // 恢复现场
		}
	}
	// 另一种情况，就是新开一辆缆车
	sum[k] = w[u] //一共k辆车，下标是从0到k-1, 所以sum[k]就是第k+1辆车
	dfs(u+1, k+1)
	sum[k] = 0 // 恢复现场
}
