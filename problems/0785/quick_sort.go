package main

import "fmt"

// https://www.acwing.com/problem/content/description/787/

func quickSort(q []int, l, r int) {
	if l >= r {
		return
	}
	// partition
	x := q[(l+r)>>1]
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
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}

func main() {
	var n int
	q := make([]int, 1e5+10)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	quickSort(q, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}
