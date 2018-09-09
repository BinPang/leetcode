package main

func main() {
	println("input:1, 3,target:2,need:-1,return:", search([]int{1, 3}, 2))
	//return
	println("input:2, 1,target:2,need:0,return:", search([]int{2, 1}, 2))
	println("input:2, 1,target:1,need:1,return:", search([]int{2, 1}, 1))
	println("input:4,5,6,7,0,1,2,target:0,need:4,return:", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	println("input:4,5,6,7,0,1,2,target:3,need:-1,return:", search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	println("input:1,target:1,need:0,return:", search([]int{1}, 1))
}

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	start := 0
	end := l - 1
	if target == nums[start] {
		return 0
	}
	if target == nums[end] {
		return l - 1
	}
	middle := 0
	for {
		if start == end {
			return -1
		}
		if end - start == 1 {
			if nums[start] == target {
				return start
			} else if nums[end] == target {
				return end
			} else {
				return -1
			}
		}
		middle = (start + end) / 2
		//println(start, end, middle)
		if target == nums[middle] {
			return middle
		}
		if nums[middle] > nums[start] {
			if target < nums[middle] && target > nums[start] {
				end = middle
			} else {
				start = middle
			}
		} else {
			if target > nums[middle] && target < nums[end] {
				start = middle
			} else {
				end = middle
			}
		}
	}
	return -1
}
