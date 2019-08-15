package main

import "fmt"
//why some one is so clever

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	println(fmt.Sprintf("%+v", obj.d))
	println(obj.SumRange(0, 2))
	println(obj.SumRange(2, 5))
	println(obj.SumRange(0, 5))
}

type NumArray struct {
	d []int
}

func Constructor(nums []int) NumArray {
	l := len(nums)
	for i := 1; i < l; i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{d: nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 || i > j || j >= len(this.d) {
		return 0
	}
	if i == 0 {
		return this.d[j]
	}
	return this.d[j] - this.d[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
