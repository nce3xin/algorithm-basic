package floatbinarysearch

// 浮点二分以开方为例

func sqrt(x float64) float64 {
	var l, r float64 = 0, x
	// 如果结果要求保留4位小数，那这儿就写1e-6。这儿有效位数至少要多2位，经验值。
	for r-l > 1e-6 {
		mid := l + (r-l)/2
		if mid*mid > x {
			// 因为是浮点数，不需要处理+1，-1, 就很舒服
			r = mid
		} else {
			l = mid
		}
	}
	return l
}
