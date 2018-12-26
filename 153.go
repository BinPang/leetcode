package main

func main() {
	println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	println(findMin([]int{4, 5, 6, 7, 8, 9, 2}))
	println(findMin([]int{7, 0, 1, 2, 3, 4, 5, 6}))
}

/**
Input: [3,4,5,1,2]
Output: 1
Input: [4,5,6,7,0,1,2]
Output: 0
 */
func findMin(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	}
	start := 0
	end := l - 1
	middle := 0
	for {
		//println(start, end)
		if start+1 == end {
			if nums[start] > nums[end] {
				return nums[end]
			} else {
				return nums[start]
			}
		}
		middle = (start + end) >> 1
		if nums[start] > nums[end] {
			if nums[middle] > nums[start] {
				start = middle
			} else {
				end = middle
			}
		} else {
			return nums[start]
		}
	}
}
