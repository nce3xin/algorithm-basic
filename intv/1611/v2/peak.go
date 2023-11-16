package v1

// Forward declaration of query API.
func query(x int) int {
	// ...
	// return int means nums[x].
	return x // just an example, not the correct return value
}

// 这种写法无法AC，我很疑惑为什么这种写法不行
func findPeakElement(N int) int {
	l, r := 0, N-1
	for l < r {
		mid := (l + r + 1) >> 1
		if query(mid) < query(mid+1) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
