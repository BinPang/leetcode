package main

import "fmt"

func main() {
	println(fmt.Sprintf("%+v", productExceptSelf([]int{3, 5})))
	println(fmt.Sprintf("%+v", productExceptSelf([]int{3, 5, 6, 2, 4})))
}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	r := make([]int, l)
	r[0] = nums[0]
	for i := 1; i < l-1; i++ {
		r[i] = nums[i] * r[i-1]
	}
	tmp := 1
	for i := l - 1; i > 0; i-- {
		r[i] = r[i-1] * tmp
		tmp *= nums[i]
	}
	r[0] = tmp

	return r
}
