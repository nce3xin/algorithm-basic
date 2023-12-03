package main

func NumberOf1(n int) int {
	res := 0
	for n != 0 {
		n -= lowBit(n)
		res++
	}
	return res
}

func lowBit(n int) int {
	return n & -n
}
