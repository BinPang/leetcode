package main

import "strconv"

func main() {
	println(getPermutation(4, 9))
	println(getPermutation(4, 8))
	println(getPermutation(8, 678))
	println(getPermutation(4, 10))
	println(getPermutation(4, 24))
}

func getPermutation(n int, k int) string {
	//Given n will be between 1 and 9 inclusive.
	r := make([]int, 0)
	per := make([]int, n)
	num := make([]int, n)
	per[0] = 1
	num[0] = 1
	for i := 1; i < n; i++ {
		per[i] = (i+1) * per[i-1]
		num[i] = i + 1
	}
	k--//start from 1 to zero
	for i := n-1; i > 0; i-- {
		loc := k / per[i-1]
		r = append(r, num[loc])
		k = k - loc*per[i-1]
		num = append(num[0:loc], num[loc+1:]...)
	}
	r = append(r, num[0])
	r0 := 0
	for i := 0; i < n; i++ {
		r0 = r0 * 10 + r[i]
	}

	return strconv.Itoa(r0)
}
