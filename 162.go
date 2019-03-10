package main

func main() {
	println(findPeakElement([]int{1, 2, 3}))

	println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	println(findPeakElement([]int{1, 2, 1}))
	println(findPeakElement([]int{2, 1, 2}))
}

func findPeakElement(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 0
	}
	start := 0
	end := l - 1
	middle := 0
	for {
		middle = (start + end) / 2
		//println(fmt.Sprintf("%d,%d,%d,%+v", start, end, middle, nums))
		if middle == 0 {
			if nums[0] > nums[1] {
				return 0
			} else {
				return 1
			}
		}
		if start+1 == end {
			if nums[start] > nums[end] {
				return start
			} else {
				return end
			}
		}
		if nums[middle] > nums[middle-1 ] && nums[middle] > nums[middle+1 ] {
			return middle
		}
		if nums[middle] > nums[start] {
			if nums[middle-1] < nums[middle] {
				start = middle
			} else {
				end = middle
			}
		} else {
			end = middle
		}
	}
	return 0
}
