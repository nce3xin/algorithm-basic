package main

import "fmt"

// 时间复杂度O(n)，因为快速选择算法，每次选择一半的区间，
// 所以时间复杂度 = n + n/2 + n/4 + n/8 ... = n(1+1/2+1/4+1/8+...) <= 2n
// 如果用排序，时间复杂度就是O(n*log n)了

const N int = 1e5 + 10

var q []int = make([]int, N)

func main() {
	var n, k int
	fmt.Scanf("%d%d", &n, &k)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	kth := quickSort(0, n-1, k)
	fmt.Printf("%d", kth)
}

func quickSort(l, r, k int) int {
	// recursive termination condition
	// 当前函数满足k是一定在[l,r]区间内的，如果l==r，说明只有一个元素，就是k
	if l == r {
		return q[l]
	}
	// pivot
	x := q[(l+r)>>1]
	// partition
	i, j := l-1, r+1
	for i < j {
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
	// 左半边元素个数
	sl := j - l + 1
	// 快速选择算法，根据k和左半边元素个数的比较，选择一边进行递归，不用递归两边
	if k <= sl {
		// 递归左边
		return quickSort(l, j, k)
	}
	// 第k大的数在右半边是第k-sl个数
	return quickSort(j+1, r, k-sl)
}
