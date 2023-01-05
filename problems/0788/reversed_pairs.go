package main

import "fmt"

const N int = 1e5 + 10

var tmp []int = make([]int, N)

func main() {
	var n int
	q := make([]int, N)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	res := mergeSort(q, 0, n-1)
	fmt.Printf("%d", res)
}

// return number of reversed pairs
func mergeSort(q []int, l, r int) int64 {
	if l >= r {
		return 0
	}
	mid := (l + r) >> 1
	res := mergeSort(q, l, mid) + mergeSort(q, mid+1, r)
	// merge
	k := 0 // 已排好序的元素个数
	i, j := l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			k++
			i++
		} else {
			tmp[k] = q[j]
			k++
			j++
			res += int64(mid - i + 1)
		}
	}
	// 扫尾
	for i <= mid {
		tmp[k] = q[i]
		i++
		k++
	}
	for j <= r {
		tmp[k] = q[j]
		j++
		k++
	}
	// 物归原主
	for i, j = l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = tmp[j]
	}
	return res
}
