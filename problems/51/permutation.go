package main

import "sort"

const N int = 10

var st [N]bool
var path []int
var res [][]int

func permutation(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	dfs(0, n, nums)
	return res
}

// 枚举顺序：每个位置放哪个数
func dfs(u int, n int, nums []int) {
	if u == n {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if !st[i] {
			st[i] = true
			path = append(path, nums[i])
			dfs(u+1, n, nums)
			st[i] = false
			path = path[:len(path)-1]
			// 去重
			// 这儿是for，别写成if了
			for i+1 < n && nums[i] == nums[i+1] {
				i++
			}
		}
	}
}
