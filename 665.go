package main

func main() {
	println(checkPossibility([]int{3, 4, 2, 3}))
}

func checkPossibility(nums []int) bool {
	l := len(nums)
	if l <= 2 {
		return true
	}
	isWell := false
	if nums[0] > nums[1] {
		isWell = true
		nums[0] = nums[1]
	}
	index := 1
	for index < l {
		if nums[index] >= nums[index-1] {
			index++
			continue
		}
		if isWell {
			return false
		}
		//nums[index] < nums[index-1] //* 10 6
		if nums[index] >= nums[index-2] { // * 5 10 6
			isWell = true
			nums[index-1] = nums[index-2]
		} else { //* 8 10 6
			isWell = true
			nums[index] = nums[index-1]
		}
		index++
	}
	return true
}
