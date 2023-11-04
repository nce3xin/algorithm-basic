package main

import "fmt"

func main() {
	var a, n int
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		fmt.Scanf("%d", &a)
		if isPrime(a) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func isPrime(n int) bool {
	// 2是最小的质数，如果n小于2，那n肯定就不是质数
	if n < 2 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
