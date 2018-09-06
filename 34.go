package main

func main() {
	t3 := searchRange([]int{1}, 1)
	_printArray(t3)

	t0 := searchRange([]int{1, 2, 2, 2, 2, 2}, 2)
	_printArray(t0)

	t1 := searchRange([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 1)
	_printArray(t1)

	t2 := searchRange([]int{}, 1)
	_printArray(t2)
}

func searchRange(nums []int, target int) []int {
	l := len(nums)
	rs, re := -1, -1
	if l == 0 {
		return []int{rs, re}
	}
	start, end := 0, l-1
	middle := -1
	for {
		middle = (start + end) / 2
		if nums[middle] == target {
			rs = middle
			re = middle
			break
		} else if nums[middle] <= target {
			start = middle + 1
		} else {
			end = middle - 1
		}
		if start > end {
			return []int{-1, -1}
		}
	}
	//println(middle)
	//find left
	start, end = 0, middle
	m1 := 0
	for {
		//println("left:", start, end, m1)
		if start+1 >= end {
			if nums[start] == target {
				rs = start
			} else {
				rs = end
			}
			break
		}
		m1 = (start + end) / 2
		if nums[m1] < target {
			start = m1
		} else {
			end = m1
		}
	}
	//println(m1)

	//find right
	start, end = middle, l-1
	m2 := 0
	for {
		//println("right:", start, end, m1)
		if start+1 >= end {
			if nums[end] == target {
				re = end
			} else {
				re = start
			}
			break
		}
		m2 = (start + end) / 2
		if nums[m2] > target {
			end = m2
		} else {
			start = m2
		}
	}
	//println(m2)
	return []int{rs, re}
}

func _printArray(t []int) {
	for _, v := range t {
		print(v, ",")
	}
	println("")
}
