package main

import "fmt"

const N int = 1e5 + 10

func main() {
	var n int
	var a []int = make([]int, N)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	r := 0 // 结果
	// s[] 存储a数组中每个元素出现的次数
	s := make([]int, N)
	// i是右边指针，j是左边指针，i>j
	for i, j := 0, 0; i < n; i++ {
		s[a[i]]++
		// a[i] 重复
		for s[a[i]] > 1 {
			s[a[j]]--
			j++
		}
		r = max(r, i-j+1)
	}
	fmt.Printf("%d", r)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
