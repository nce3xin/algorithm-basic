package main

import (
	"fmt"
	"sort"
)

const N int = 30

var n, m int
var st [N]bool

func main() {
	fmt.Scanf("%d%d", &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	// 从小到大排序
	sort.Ints(a)
	dfs(0, 0, a)
}

// s是当前已经选了几个数
func dfs(u int, s int, a []int) {
	if s > m {
		return
	}
	if s == m {
		for i := 0; i < u; i++ { // 注意这里是i<u
			if st[i] {
				fmt.Printf("%d ", a[i])
			}
		}
		fmt.Println()
		return
	}
	// 已经看完了所有数，还没选够m个
	if u == n {
		return
	}
	// 枚举当前种元素选几个，因此需要先把这段数找出来
	// 注意，此题上1572不一样，1572不要求字典序，但这题要求字典序。所以应该是优先选择前面的数（看递归搜索树）
	// 所以前面的数选的个数，应该是从大到小枚举，也就是尽可能选的多一些。最好全选。
	k := u
	for k < n && a[k] == a[u] { // 先全选
		st[k] = true
		k++
		s++
	}
	dfs(k, s, a) // 当前这段全选
	// 依次枚举当前段内选几个，个数从大到小枚举
	for i := k - 1; i >= u; i-- {
		// 去掉最后一个
		st[i] = false
		s--
		dfs(k, s, a) // 下一层
	}
}
