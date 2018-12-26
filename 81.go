package main

func main() {
	println(search([]int{0, 0, 1, 1, 2, 0}, 2))
	return
	println(search([]int{1, 1, 3}, 3))
	println(search([]int{1}, 0))
	println(search([]int{1, 1, 1, 1, 1}, 1))
	println(search([]int{1, 3}, 3))
	println(search([]int{1, 1}, 0))
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 2))
	println(search([]int{2, 5, 6, 0, 0, 1, 1}, 2))
	println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

func search(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1
	if end == -1 {
		return false
	}
	middle := 0
	for {
		println(start, end, (start+end)/2)
		middle = (start + end) / 2
		if nums[middle] == target {
			return true
		}
		if start == middle || end == middle {
			if nums[start] == target || nums[end] == target {
				return true
			} else {
				return false
			}
		}
		if nums[start] == nums[end] && nums[start] == nums[middle] {
			start++
			end--
			continue
		}
		if nums[start] < nums[middle] {
			if nums[start] <= target && target <= nums[middle] {
				end = middle
			} else {
				start = middle
			}
		} else {
			if nums[middle] <= target && target <= nums[end] {
				start = middle
			} else {
				end = middle
			}
		}
	}
}
