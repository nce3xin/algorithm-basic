package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 10

var n int
var q [N]int

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &q[i])
	}
	quickSort(0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}

func quickSort(l, r int) {
	if l >= r {
		return
	}
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
	quickSort(l, j)
	quickSort(j+1, r)
}
