package main

import "fmt"

func main() {
	// input
	var n int
	var m int
	fmt.Scanf("%d%d", &n, &m)
	q := make([]int, 1e5+10)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	var x int // query
	for ; m > 0; m-- {
		fmt.Scanf("%d", &x)
		lb := binarySearchLowerBound(q, x, 0, n-1)
		ub := binarySearchUpperBound(q, x, 0, n-1)
		fmt.Printf("%d %d\n", lb, ub)
	}
}

func binarySearchLowerBound(q []int, x int, l, r int) int {
	// 定义性质为>=x, 题目说数组是有序的，所以数组可以分成两半，右边满足性质>=x，左边不满足。
	// 所以可以通过二分寻找性质>=x的边界，也就是x的起始位置。
	for l < r {
		mid := (l + r) >> 1
		// 这里的check条件可以跟性质保持一致，性质是>=x，check也写>=x
		if q[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 二分找到的是性质>=x的边界，如果数组中不存在等于x的值，那么找到的这个边界必然>x
	// 也就是不存在
	if q[l] != x {
		return -1
	}
	// 返回l或r都可以，因为for循环结束后，l == r
	return l
}

func binarySearchUpperBound(q []int, x int, l, r int) int {
	// 定义性质为<=x, 题目说数组是有序的，所以数组可以分成两半，左边满足性质<=x，右边不满足。
	// 所以可以通过二分寻找<=x的边界，也就是x的终止位置
	for l < r {
		// 先不管三七二十一，写上mid:=(l+r)>>1，是否需要加1，后面再看
		mid := (l + r + 1) >> 1
		if q[mid] <= x {
			l = mid
		} else {
			r = mid - 1
		}
	}
	// 下面的r也可以换成l，因为for循环结束后，l == r
	if q[r] != x {
		return -1
	}
	return r
}
