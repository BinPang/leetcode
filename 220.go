package main

func main() {
	println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
	println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	tmp := 0
	for i, v := range nums {
		for j := i + 1; j < i+k+1; j++ {
			if j == len(nums) {
				break
			}
			tmp = nums[j] - v
			if tmp < 0 {
				tmp = -tmp
			}
			if tmp <= t {
				return true
			}
		}
	}
	return false
}
