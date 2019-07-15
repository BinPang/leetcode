package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})))
}

func findDuplicates(nums []int) []int {
	l := len(nums)
	r := make([]int, 0)
	v := 0
	tmp := 0
	for i := 0; i < l; i++ {
		if nums[i] < 0 {
			continue
		}
		v = nums[i]
		nums[i] = 0
		for {
			if nums[v-1] <= 0 {
				nums[v-1] -= 1
				break
			}
			tmp = nums[v-1]
			nums[v-1] = -1
			v = tmp
		}
	}
	for i := 0; i < l; i++ {
		if nums[i] == -2 {
			r = append(r, i+1)
		}
	}
	return r
}
