package main

func minMoves(nums []int) int {
	l := len(nums)
	if l == 0 {
		panic("empty")
	}
	min := nums[0] //not empty
	total := nums[0]
	for i := 1; i < l; i++ {
		total += nums[i]
		if min > nums[i] {
			min = nums[i]
		}
	}
	return total - min*l
}
