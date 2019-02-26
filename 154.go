package main

func findMin(nums []int) int {
	start := 0
	end := len(nums)-1
	r := nums[0]
	for {
		if nums[start] == nums[end] {
			start++
			end--
			continue
		}

	}

	return r
}
