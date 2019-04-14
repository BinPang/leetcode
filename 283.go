package main

func moveZeroes(nums []int) {
	zero := 0
	index := 0
	e := 0
	for e = range nums {
		if nums[e] == 0 {
			zero++
		} else {
			if zero != 0 {
				nums[index] = nums[e]
			}
			index++
		}
	}
	for zero > 0 {
		nums[index] = 0
		index++
		zero--
	}
}
