package main

func main() {
	println("true:", search([]int{3,1,2,2,2}, 1))
	return
	println("true:", search([]int{1, 2, 2, 2, 0, 1, 1}, 0))
	println("true:", search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	//return
	println("true:", search([]int{1, 3, 1, 1, 1}, 3))
	println("true", search([]int{2, 5, 6, 0, 0, 1, 2}, 2))
	println("true", search([]int{0, 0, 1, 1, 2, 0}, 2))
	println("true", search([]int{1, 1, 3}, 3))
	println("false", search([]int{1}, 0))
	println("true", search([]int{1, 1, 1, 1, 1}, 1))
	println("true", search([]int{1, 3}, 3))
	println("false", search([]int{1, 1}, 0))
	println("true", search([]int{2, 5, 6, 0, 0, 1, 2}, 2))
	println("true", search([]int{2, 5, 6, 0, 0, 1, 1}, 2))
	println("false", search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
}

func search(nums []int, target int) bool {
	l := len(nums)
	if l == 0 {
		return false
	}
	start := 0
	end := l - 1
	middle := 0
	for {
		println(start, end, (start+end)/2)
		middle = (start + end) / 2
		if start > end {
			return false
		}
		if nums[start] == target || nums[end] == target || nums[middle] == target {
			return true
		}
		if nums[start] == nums[end] {
			if nums[middle] == nums[start] {
				start++
				end--
			} else {
				if target > nums[middle] {
					start = middle + 1
				} else {
					if nums[start] < target {
						end = middle - 1
					} else {
						start = middle + 1
					}
				}
			}
		} else {
			//3,1,2,2,2
			//1
			if target > nums[middle] {
				start = middle + 1
			} else {
				if nums[start] < target {
					end = middle - 1
				} else {
					start = middle + 1
				}
			}
		}
	}
}
