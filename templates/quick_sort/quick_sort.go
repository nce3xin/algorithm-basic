package quick_sort

func quickSort(q []int, l, r int) {
	// 只有1个数或者没有数，不用排了，直接返回
	if l >= r {
		return
	}
	// partition
	x := q[(l+r)>>1] // pivot，取中间值，不能取q[l]了，否则过不了oj的最新数据
	i, j := l-1, r+1 // 由于下面do-while中要先+1, 这里i,j要先往外扩一个
	for i < j {
		// 模拟C++模板写法中的do-while
		i++
		for q[i] < x {
			i++
		}
		j--
		for q[j] > x {
			j--
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}
	// [l,j] 和 [j+1,r] 的 pivot(line14) 不能选择 q[r]，会有边界问题，导致死循环，选择数组中任意其他位置都可以。
	// [l,i] 区间若合法则其中所有值 <= pivot
	// [j,r] 区间若合法则其中所有值 >= pivot
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}
