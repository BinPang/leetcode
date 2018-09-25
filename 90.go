package main

import (
	"learn/util"
	"sort"
)

func main() {
	t0 := subsetsWithDup([]int{2, 2, 2, 1, 2})
	util.PrintArrayArray(t0)
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) //why use sort, i don't use sort, but not accepted.@TODO not sort if not the first time
	r := make([][]int, 1)
	r[0] = make([]int, 0)
	dup := map[int]bool{}
	l := len(nums)
	if l == 0 {
		return nil
	}
	for i, v := range nums {
		if dup[v] {
			continue
		} else {
			dup[v] = true
		}
		tmp := subsetsWithDup(nums[i+1:])
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
