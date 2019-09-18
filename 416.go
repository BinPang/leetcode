package main

import "sort"

func canPartition(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	psum := 0
	for _, v := range nums {
		psum += v

	}
}
