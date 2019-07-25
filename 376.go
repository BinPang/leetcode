package main

func main() {
	println(wiggleMaxLength([]int{1}))
	println(wiggleMaxLength([]int{0, 0}))
	println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	println(wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

//no dp
//1, 17, 5, 10, 13, 15, 10, 5, 16, 8
// 16, -12, 5, 3, 2, -5, -5, 11, -8
//combine 5,3,2 and combine -5,-5
func wiggleMaxLength(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	up := true
	index := 1
	r := 1
	for {
		if index == l {
			break
		}
		if nums[index] == nums[index-1] {
			index++
			continue
		}
		if nums[index] < nums[index-1] {
			up = false
		}
		r += 1
		break
	}
	for index < l {
		if (nums[index] > nums[index-1] && !up) || (nums[index] < nums[index-1] && up) {
			r += 1
			up = !up
		}
		index++
	}
	return r
}
