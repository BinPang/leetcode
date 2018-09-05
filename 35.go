package main

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	if target <= nums[0] {
		return 0
	}
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return l
}
