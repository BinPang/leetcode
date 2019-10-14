package main

func main() {
	println(findUnsortedSubarray([]int{1, 2, 4, 5, 3}))
	return
	println(findUnsortedSubarray([]int{2, 2, 1, 3, 4, 5}))
	println(findUnsortedSubarray([]int{1, 3, 2, 2, 2}))
	println(findUnsortedSubarray([]int{2, 1}))
	println(findUnsortedSubarray([]int{1, 2, 3}))
	println(findUnsortedSubarray([]int{1, 3, 2, 4, 5, 6}))
	println(findUnsortedSubarray([]int{1, 3, 2, 1, 5, 6}))
	println(findUnsortedSubarray([]int{1, 3, 2, 5, 4, 6}))
}

func findUnsortedSubarray(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}
	start := -1
	dup := 1
	for i := 1; i < l; i++ {
		if nums[i] < nums[i-1] {
			start = i - dup
			break
		}
		if nums[i] == nums[i-1] {
			dup++
		} else {
			dup = 1
		}
	}
	if start < 0 {
		return 0
	}
	end := -1
	dup = 1
	for i := l - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			end = i + dup
			break
		}
		if nums[i] == nums[i+1] {
			dup++
		} else {
			dup = 1
		}
	}
	return end - start + 1
}
