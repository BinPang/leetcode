package main

/*
Input: [3,0,1]
Output: 2

Input: [9,6,4,2,3,5,7,0,1]
Output: 8
*/
func main() {
	println(missingNumber([]int{3, 0, 1}))
	println(missingNumber([]int{2, 0, 1}))
	println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	for _, v := range nums {
		if v == n || v == -1 {
			continue
		}
		newIndex := v
		oldIndex := 0
		for {
			if newIndex == n || newIndex == -1 {
				break
			}
			oldIndex = newIndex
			newIndex = nums[newIndex]
			nums[oldIndex] = -1
		}
	}
	for k, v := range nums {
		if v != -1 {
			return k
		}
	}
	return n
}
