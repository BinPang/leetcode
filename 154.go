package main

func main() {
	println(findMin([]int{2, 2, 2, 0, 1}))
	println(findMin([]int{1, 3, 5}))
	println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	middle := 0
	for {
		middle = (start + end) / 2
		if start == middle {
			if nums[start] > nums[end] {
				return nums[end]
			} else {
				return nums[start]
			}
		}
		if nums[start] == nums[end] {
			if nums[middle] == nums[start] {
				start++
				end--
			} else if nums[middle] < nums[start] {
				end = middle
			} else {
				start = middle
			}
		} else if nums[start] < nums[end] {
			if nums[middle] < nums[start] {
				//not exist
			} else if nums[middle] == nums[start] {
				return nums[start]
			} else if nums[middle] < nums[end] {
				return nums[start]
			} else if nums[middle] == nums[end] {
				end = middle
			} else {
				//not exist
			}
		} else { //start > end
			if nums[middle] <= nums[end] {
				end = middle
			} else if nums[middle] < nums[start] {
				//not exist
			} else if nums[middle] >= nums[start] {
				start = middle
			} else {
				//not exist
			}
		}
	}
}
