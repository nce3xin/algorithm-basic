package main

import "fmt"

func main() {
	var n float64
	fmt.Scanf("%f", &n)
	// 性质：x*x*x > n
	var l, r float64
	// 注意l，r的取值，不能0, n
	// [l,r]区间只要包含答案就可以，[-10000,10000]的三次方根，也在[-10000,10000]
	l, r = -1e4, 1e4
	for r-l > 1e-8 {
		mid := (l + r) / 2
		if mid*mid*mid > n {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Printf("%.6f", l)
}
