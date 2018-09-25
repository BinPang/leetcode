package main

import "learn/util"

func main() {
	t0 := subsets([]int{1, 2, 3})
	util.PrintArrayArray(t0)
}

func subsets(nums []int) [][]int {
	r := make([][]int, 1)
	r[0] = make([]int, 0)
	l := len(nums)
	if l == 0 {
		return nil
	}
	for i, v := range nums {
		tmp := subsets(nums[i+1:])
		if tmp != nil {
			for _, v1 := range tmp {
				r = append(r, append([]int{v}, v1...))
			}
		} else {
			r = append(r, []int{v})
		}
	}
	return r
}
