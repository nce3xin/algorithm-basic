package main

import "fmt"

var tmp []int = make([]int, 1e5+10)

func mergeSort(q []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)
	k := 0
	i, j := l, mid+1
	for i <= mid && j <= r {
		if q[i] < q[j] {
			tmp[k] = q[i]
			k++
			i++
		} else {
			tmp[k] = q[j]
			k++
			j++
		}
	}
	for i <= mid {
		tmp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		tmp[k] = q[j]
		k++
		j++
	}
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = tmp[j]
	}
}

func main() {
	var n int
	q := make([]int, 1e6+10)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	mergeSort(q, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}
