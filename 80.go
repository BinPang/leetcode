package main

func main() {
	removeDuplicates([]int{1, 1, 1, 2, 2, 3})
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}
	r := 2
	one := nums[0]
	two := nums[1]
	for i := 2; i < l; i++ {
		if nums[i] == one {
			continue
		}
		one = two
		two = nums[i]
		nums[r] = nums[i]
		r++
	}
	return r
}
