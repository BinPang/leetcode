package main

func containsNearbyDuplicate(nums []int, k int) bool {
	tmp := make(map[int]int, len(nums))
	for i, v := range nums {
		if step, ok := tmp[v]; ok {
			if i-step <= k {
				return true
			}
		}
		tmp[v] = i
	}
	return false
}
