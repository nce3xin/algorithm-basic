package main

func maxSubArray(nums []int) int {
	res := int(-2e5)
	// s表示以当前数的前一个数结尾的子数组的和的最大值
	s := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		x := nums[i]
		if s <= 0 {
			s = 0
		}
		s = s + x
		res = max(res, s)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
