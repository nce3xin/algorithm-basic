package main

import (
	"fmt"
	"sort"
)

const N int = 20

var n int

var st [N]bool // bool数组表示每个数有没有选

func main() {
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i]) // a[i]表示原数组中的每个数
	}
	// a[]从小到大排序
	sort.Ints(a)
	dfs(0, a)
}

// u 表示的不是第几段，而是每一段的第一个数在整个a数组里的下标
func dfs(u int, a []int) {
	// 因为main函数中dfs(0),不是dfs(1)，所以这里判断条件是u==n。否则是u > n
	if u == n {
		// 枚举a[]的每个数，有没有被选过，被选的就输出
		for i := 0; i < n; i++ {
			if st[i] {
				fmt.Printf("%d ", a[i])
			}
		}
		fmt.Println()
		return
	}
	// 枚举当前种元素选几个，因此需要先把这段数找出来
	// 怎么找呢？有一个非常固定的写法。下面这两行代码特别特别常见，一点更牢记
	k := u
	for k < n && a[k] == a[u] {
		// k一直往后走，k要么走到边界，要么走到下一段的开头
		k++
	}
	// 需要先枚举一下选0个的情况
	// 一个都不标记成true
	dfs(k, a)
	// 枚举选前i个
	for i := u; i < k; i++ {
		st[i] = true // 选前i个，所以把前i个的st[i]都设为true
		// dfs到下一层，下一层应该是从下一段的开头开始枚举。下一层的开头的下标是k
		dfs(k, a)
		// 这儿千万别写st[i] = false了，不是不做恢复现场，而是不应该在这儿恢复现场
		// 因为第i个分支表示选前i个，也就是要把[u,i]全部赋成true。也就是第i个分支要求[u,i]全部是true
		// 如果这儿写了st[i] = false, 就表示只将第i个置成true, 但其实应该把[u,i]全部置成true
		// st[i] = false
	}
	// 整个分支搜完之后，整个分支都要还原
	// 恢复现场
	for i := u; i < k; i++ {
		st[i] = false
	}
}
