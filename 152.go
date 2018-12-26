package main

/**
Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 */

func main() {
	println(maxProduct([]int{2, -5, -2, -4, 3}))
	return
	println(maxProduct([]int{3, -1, 4}))
	println(maxProduct([]int{-2, 0, -1}))
	println(maxProduct([]int{2, 3, -2, 4}))
	println(maxProduct([]int{0, 3, 2, 4}))
	println(maxProduct([]int{-2, 3, -2, 4}))
	println(maxProduct([]int{-2, 3, 0, 0, 0, -2, 4, -2, 0, 0, 3, 4}))
}

func maxProduct(nums []int) int {
	r := nums[0]
	l := len(nums)
	max := r
	min := r
	for i := 1; i < l; i++ {
		if nums[i] == 0 {
			if max < 0 {
				max = 0
			}
		}
	}
	return r
}
