package main

const INF = 1e9

func getMinimumValue(N int) []int {
	// 二分列
	l, r := 0, N-1
	for l < r {
		mid := (l + r) >> 1
		// 求mid列上的最小值
		var k int       // 最小值所在的行
		val := int(INF) // 中间列的最小值
		for i := 0; i < N; i++ {
			t := query(i, mid)
			if t < val {
				val = t
				k = i
			}
		}
		// 求最小值左右的值
		var left, right int
		// 判断mid是不是最左边的列，边界条件
		if mid > 0 {
			left = query(k, mid-1)
		} else {
			left = INF
		}
		// 判断mid是不是最右边的列，边界条件
		if mid < N-1 {
			right = query(k, mid+1)
		} else {
			right = INF
		}
		// 中间列的最小值，小于其左右两边的值，说明找到了一个极小值
		if val < left && val < right {
			// 输出极小值所在的行和列
			return []int{k, mid}
		}
		if left < val { // 二分左边区域，朝小于自己的方向走
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 走到这里的时候，l等于r，这时候范围就剩一列了，再求这列的最小值，就是极小值
	var k int       // 这列的最小值所在行
	val := int(INF) // 这列的最小值
	for i := 0; i < N; i++ {
		t := query(i, r) // 这里写query(i, l)也行，因为这时l和r相等
		if t < val {
			val = t
			k = i
		}
	}
	return []int{k, r}
}

// Forward declaration of query API.
// return int means matrix[x][y].
func query(a int, b int) int {
	// ...
	return 0
}
