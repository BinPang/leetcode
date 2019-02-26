package main

func main() {
	println("false", search([]int{1, 3}, 2))
	println("false", search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	println("true:", search([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 25))
	println("true:", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	println("true:", search([]int{0, 1, 1, 2, 0, 0}, 2))
	println("true:", search([]int{0, 0, 1, 1, 2, 0}, 2))
	println("true:", search([]int{3, 1, 2, 2, 2}, 1))
	println("true:", search([]int{1, 2, 2, 2, 0, 1, 1}, 0))
	println("true:", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	println("true:", search([]int{1, 3, 1, 1, 1}, 3))
	println("true", search([]int{2, 5, 6, 0, 0, 1, 2}, 2))
	println("true", search([]int{1, 1, 3}, 3))
	println("false", search([]int{1}, 0))
	println("true", search([]int{1, 1, 1, 1, 1}, 1))
	println("true", search([]int{1, 3}, 3))
	println("false", search([]int{1, 1}, 0))
	println("true", search([]int{2, 5, 6, 0, 0, 1, 2}, 2))
	println("true", search([]int{2, 5, 6, 0, 0, 1, 1}, 2))
}

func search(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	start := 0
	end := l - 1
	middle := 0
	//println(fmt.Sprintf("%+v,%d", nums, target))
	for {
		middle = (start + end) / 2
		//println(start, end, middle)
		if start > l-1 || end < 0 || start > end {
			return false
		}
		if nums[start] == target || nums[end] == target || nums[middle] == target {
			return true
		}
		if nums[start] == nums[end] {
			start++
			end--
			continue
		}
		if start < l-1 && nums[start] == nums[start+1] {
			start++
			continue
		}
		if end > 0 && nums[end] == nums[end-1] {
			end--
			continue
		}
		if start+1 >= end {
			return false
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
}
