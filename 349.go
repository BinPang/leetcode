package main

func intersection(nums1 []int, nums2 []int) []int {
	tmp := make(map[int]int, len(nums1))
	r := make([]int, 0)
	for _, v := range nums1 {
		tmp[v] = 1
	}
	for _, v := range nums2 {
		if tmp[v] == 1 {
			r = append(r, v)
			tmp[v] = 2
		}
	}
	return r
}
