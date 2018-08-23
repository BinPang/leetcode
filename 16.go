package main

import "sort"

func main() {

	println("input:[1,1,-1,-1,3] and -1,return:", threeSumClosest([]int{1, 1, -1, -1, 3}, -1))
	println("input:[-1, 2, 1, -4] and 1,return:", threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	if l < 3 {
		panic("at least 3 items")
	}
	r := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)

	j := 0
	k := 0
	v := 0
	tmp1 := 0
	tmp2 := 0
	for i := 0; i < l-2; i++ {
		j = i + 1
		k = l- 1
		for j < k {
			v = nums[i] + nums[j] + nums[k]

			if v == target {
				return v
			}
			if r > target {
				tmp1 = r - target
			} else {
				tmp1 = target-r
			}
			if v > target {
				tmp2 = v-target
			} else {
				tmp2 = target-v
			}
			if tmp2 < tmp1 {
				r = v
			}
			if v > target {
				k--
			} else if v < target {
				j++
			} else {
				return target
			}
		}
	}
	return r
}
