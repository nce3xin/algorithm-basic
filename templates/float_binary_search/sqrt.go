package floatbinarysearch

// 浮点二分以开方为例

func sqrt(x float64) float64 {
	// 这里容易出bug
	// r不能是x，因为如果x<1,比如x=0.01，开方应该是0.1。而0.1不在[0,x]范围内
	// 严谨来说，r = max(1,x)
	var l, r float64 = 0, max(1, x)
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

func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
