package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})))
	println(fmt.Sprintf("%+v", nextGreaterElement([]int{4, 1, 2}, []int{3, 2, 1, 4})))
	println(fmt.Sprintf("%+v", nextGreaterElement([]int{4, 1, 2}, []int{2, 1, 3, 4})))
	println(fmt.Sprintf("%+v", nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return make([]int, 0)
	}
	l := len(nums2)
	r := make([]int, len(nums1))
	tmp := make(map[int]int, l)
	need := make([]int, 0)
	needIndex := -1
	for _, v := range nums2 {
		for {
			if needIndex < 0 {
				break
			}
			if need[needIndex] < v {
				tmp[need[needIndex]] = v
			} else {
				break
			}
			needIndex--
		}
		if len(need) > needIndex+1 {
			need[needIndex+1] = v
		} else {
			need = append(need, v)
		}
		needIndex++
	}
	//println(fmt.Sprintf("::::%+v", tmp))
	for k, v := range nums1 {
		if v1, ok := tmp[v]; ok {
			r[k] = v1
		} else {
			r[k] = -1
		}

	}
	return r
}
