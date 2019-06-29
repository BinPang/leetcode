package main

import "fmt"

func main() {
	println(1 ^ 0)
	println(3 ^ 1)
	println(fmt.Sprintf("%+v", singleNumber([]int{1, 2, 1, 3, 2, 5})))
}

func singleNumber(nums []int) []int {
	l := len(nums)
	if l == 0 {
		return nil
	}
	tmp := nums[0]
	for i := 1; i < l; i++ {
		tmp ^= nums[i]
	}
	//find 1 location
	i := 0
	for {
		if tmp&1 == 1 {
			break
		}
		i++
		tmp >>= 1
	}
	iValue := 1 << uint(i)
	num1, num2 := 0, 0
	init1, init2 := false, false
	for i = 0; i < l; i++ {
		if nums[i]&iValue == 0 {
			if init1 {
				num1 = nums[i]
				init1 = true
			} else {
				num1 ^= nums[i]
			}
		} else {
			if init2 {
				num2 = nums[i]
				init2 = true
			} else {
				num2 ^= nums[i]
			}
		}
	}
	return []int{num1, num2}
}
