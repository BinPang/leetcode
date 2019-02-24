package main

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]bool, len(nums))
	for _, v := range nums {
		if tmp[v] {
			return true
		}
		tmp[v] = true
	}
	return false
}
